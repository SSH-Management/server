package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/models"
)

const (
	TypeNewUser                     = "create:user"
	TypeNotifyServerForNewUser      = "notify:server"
	TypeNotifyServersForDeletedUser = "delete:user"
)

func NewUserNotification(user dto.User, publicKey string) (*asynq.Task, error) {
	bytes, err := json.Marshal(dto.CreateUser{User: user, PublicSSHKey: publicKey})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(
		TypeNewUser,
		bytes,
		asynq.MaxRetry(3),
	), nil
}

func NewNotifyServerForNewUser(server models.Server, user dto.CreateUser) (*asynq.Task, error) {
	bytes, err := json.Marshal(dto.NewUserNotification{
		User: user,
		Server: struct {
			Name      string "json:\"name,omitempty\""
			IpAddress string "json:\"ip,omitempty\""
		}{
			Name: server.Name,
			IpAddress: server.IpAddress,
		},
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(
		TypeNotifyServerForNewUser,
		bytes,
		asynq.MaxRetry(100),
	), nil
}

func NewNotifyServerForDeletedUser(server models.Server, username string) (*asynq.Task, error) {
	bytes, err := json.Marshal(dto.UserDeletedNotification{
		Username: username,
		Server: struct {
			Name      string "json:\"name,omitempty\""
			IpAddress string "json:\"ip,omitempty\""
		}{
			Name: server.Name,
			IpAddress: server.IpAddress,
		},
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(
		TypeNotifyServersForDeletedUser,
		bytes,
		asynq.MaxRetry(100),
	), nil
}
