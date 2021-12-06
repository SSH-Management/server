package processors

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/hibiken/asynq"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/SSH-Management/protobuf/client/users"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
	"github.com/SSH-Management/server/pkg/tasks"
)

var (
	_ asynq.Handler = &NewUserCreated{}
	_ asynq.Handler = &NotifyServerForNewUser{}
)

type (
	NewUserCreated struct {
		servers map[string][]models.Server
		once    sync.Once
		mutext  sync.RWMutex

		db     *gorm.DB
		queue  *asynq.Client
		logger *log.Logger
	}

	NotifyServerForNewUser struct {
		port            uint16
		mutext          sync.RWMutex
		grpcConnections map[string]*grpc.ClientConn

		logger *log.Logger
	}
)

func NewUserCreatedProcessor(db *gorm.DB, queue *asynq.Client, logger *log.Logger) *NewUserCreated {
	return &NewUserCreated{
		once:    sync.Once{},
		servers: make(map[string][]models.Server),
		mutext:  sync.RWMutex{},

		db:     db,
		queue:  queue,
		logger: logger,
	}
}

func NewNotifyServerForNewUser(logger *log.Logger) *NotifyServerForNewUser {
	return &NotifyServerForNewUser{
		port:            9999,
		mutext:          sync.RWMutex{},
		grpcConnections: make(map[string]*grpc.ClientConn),

		logger: logger,
	}
}

func (n *NewUserCreated) fetchServers(ctx context.Context) error {
	n.logger.Debug().Msg("Fetching all servers")
	n.mutext.Lock()

	var servers []models.Server

	result := n.db.
		WithContext(ctx).
		Preload("Group").
		Find(&servers)

	if result.Error != nil {
		n.logger.Error().
			Err(result.Error).
			Msg("Error while fetching servers")

		return result.Error
	}

	n.logger.Debug().
		Int("servers_legnth", len(servers)).
		Msg("Number of servers fetched from database")

	for _, server := range servers {
		if _, ok := n.servers[server.Group.Name]; !ok {
			n.servers[server.Group.Name] = make([]models.Server, 0, 10)
		}

		n.servers[server.Group.Name] = append(n.servers[server.Group.Name], server)
	}

	n.logger.Debug().Int("groups_length", len(n.servers))

	n.mutext.Unlock()
	return nil
}

func (n *NewUserCreated) ProcessTask(ctx context.Context, task *asynq.Task) error {
	if err := n.fetchServers(ctx); err != nil {
		return err
	}

	var notification dto.NewUserNotification

	payload := task.Payload()

	err := json.Unmarshal(payload, &notification)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(len(notification.Groups))

	errCh := make(chan error, len(notification.Groups))

	n.mutext.RLock()

	for _, group := range notification.Groups {
		servers, ok := n.servers[group]

		if !ok {
			continue
		}

		go func(wg *sync.WaitGroup, server []models.Server) {
			defer wg.Done()
			for _, server := range servers {
				task, err := tasks.NewNotifyServerForNewUser(
					server,
					notification.User,
					notification.PublicSSHKey,
				)
				if err != nil {
					errCh <- err
					return
				}

				_, err = n.queue.Enqueue(task)

				if err != nil {
					errCh <- err
					return
				}

				n.logger.Debug().
					Str("server", server.Name).
					Str("ip", server.IpAddress).
					Msg("Create user sent to server")
			}
		}(&wg, servers)
	}

	wg.Wait()
	n.mutext.RUnlock()
	close(errCh)

	for err := range errCh {
		if err != nil {
			n.logger.Error().
				Err(err).
				Msg("Error while sending into another queue")
		}
	}

	select {
	case <-ctx.Done():
		return errors.New("timeout...")
	default:
		return nil
	}
}

func (n *NotifyServerForNewUser) getConnectionToClient(ip string) (users.UserServiceClient, error) {
	n.mutext.RLock()
	if conn, ok := n.grpcConnections[ip]; ok {
		n.mutext.RUnlock()
		return users.NewUserServiceClient(conn), nil
	}

	n.mutext.RUnlock()
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", ip, n.port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	n.mutext.Lock()
	defer n.mutext.Unlock()

	n.grpcConnections[ip] = conn

	return users.NewUserServiceClient(conn), nil
}

func (n *NotifyServerForNewUser) ProcessTask(ctx context.Context, task *asynq.Task) error {
	payload := task.Payload()

	var notification dto.NewUserForClientsNotification

	if err := json.Unmarshal(payload, &notification); err != nil {
		return err
	}

	client, err := n.getConnectionToClient(notification.Server.IpAddress)
	if err != nil {
		return err
	}

	_, err = client.Create(ctx, &users.CreateUserRequest{
		User:      notification.User,
		PublicKey: notification.PublicSSHKey,
	})

	if err != nil {
		n.logger.Error().
			Str("server", notification.Server.IpAddress).
			Err(err).
			Msg("Error while creating new user on the client")
	}

	return nil
}

func (n *NotifyServerForNewUser) Close() error {
	n.mutext.Lock()
	defer n.mutext.Unlock()

	for server, conn := range n.grpcConnections {
		if err := conn.Close(); err != nil {
			n.logger.Error().
				Err(err).
				Str("server", server).
				Msg("Error while closing the connection")
		}
	}

	return nil
}
