package processors

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
	"github.com/SSH-Management/server/pkg/tasks"
)

var (
	_ asynq.Handler = &NewUserCreated{}
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
	}
)

func NewUserCreatedProcessor(db *gorm.DB, queue *asynq.Client, logger *log.Logger) *NewUserCreated {
	return &NewUserCreated{
		once:    sync.Once{},
		servers: make(map[string][]models.Server, 0),
		mutext:  sync.RWMutex{},

		db:     db,
		queue:  queue,
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

	var user dto.CreateUser

	payload := task.Payload()

	err := json.Unmarshal(payload, &user)

	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(len(user.User.Groups))

	errCh := make(chan error, 100)
	defer close(errCh)

	n.mutext.RLock()

	for _, group := range user.User.Groups {
		servers := n.servers[group]

		go func(wg *sync.WaitGroup, server []models.Server) {
			defer wg.Done()
			for _, server := range servers {
				task, err := tasks.NewNotifyServerForNewUser(server, user)
				if err != nil {
					errCh <- err
					return
				}

				_, err = n.queue.Enqueue(task)

				if err != nil {
					errCh <- err
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

	select {
	case err := <-errCh:
		if err != nil {
			n.logger.Error().Err(err).Msg("Error while sending into another queue")
			return err
		}
	case <-ctx.Done():
		return errors.New("timeout...")
	}

	return nil
}

func (n *NotifyServerForNewUser) ProcessTask(ctx context.Context, task *asynq.Task) error {
	payload := task.Payload()

	var notification dto.NewUserNotification

	if err := json.Unmarshal(payload, &notification); err != nil {
		return err
	}

	// TODO: Add Client SDK

	return nil
}
