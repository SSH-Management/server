package tasks

import (
	"encoding/json"
	"time"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/hibiken/asynq"
)

const (
	TypeNotifyServersForNewUser     = "create:user"
	TypeNotifyServersForDeletedUser = "delete:user"
)

func NewNotifyServerForNewUser(user dto.User, publicKey string) (*asynq.Task, error) {
	bytes, err := json.Marshal(dto.CreateUser{User: user, PublicSSHKey: publicKey})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(
		TypeNotifyServersForNewUser,
		bytes,
		asynq.MaxRetry(3),
		asynq.Timeout(200*time.Millisecond),
	), nil
}
