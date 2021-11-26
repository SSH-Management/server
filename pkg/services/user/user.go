package user

import (
	"context"
	"fmt"

	user "github.com/SSH-Management/linux-user"
	ssh "github.com/SSH-Management/ssh"
	"github.com/SSH-Management/utils/v2"
	"github.com/hibiken/asynq"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
	userepo "github.com/SSH-Management/server/pkg/repositories/user"
	"github.com/SSH-Management/server/pkg/tasks"
)

var _ Interface = &Service{}

type (
	Service struct {
		userRepo        userepo.Interface
		unixUserService user.UnixInterface

		logger *log.Logger
		queue  *asynq.Client
	}

	Interface interface {
		Get(context.Context) ([]models.User, error)
		Create(context.Context, dto.CreateUser) (models.User, []byte, error)
	}
)

func New(
	userRepo userepo.Interface,
	unixUserService user.UnixInterface,
	logger *log.Logger,
	queue *asynq.Client,
) Service {
	return Service{
		userRepo:        userRepo,
		unixUserService: unixUserService,
		logger:          logger,
		queue:           queue,
	}
}

func (s Service) Get(ctx context.Context) ([]models.User, error) {
	return s.userRepo.Find(ctx)
}

func (s Service) Create(ctx context.Context, u dto.CreateUser) (models.User, []byte, error) {
	unixUser, err := s.unixUserService.Create(ctx, u.User)
	if err != nil {
		return models.User{}, nil, err
	}

	err = ssh.AddToAuthorizedKeys(
		fmt.Sprintf("%s/.ssh/authorized_keys", unixUser.HomeFolder),
		u.PublicSSHKey,
		unixUser.UserId,
		unixUser.GroupId,
	)

	username := u.GetUser().GetUsername()

	if err != nil {
		s.deleteUserFromSystem(ctx, username)
		return models.User{}, nil, err
	}

	key, err := ssh.New(unixUser.UserId, unixUser.GroupId, unixUser.HomeFolder)
	if err != nil {
		s.deleteUserFromSystem(ctx, username)
		return models.User{}, nil, err
	}

	if err := key.Write(); err != nil {
		s.deleteUserFromSystem(ctx, username)
		return models.User{}, nil, err
	}

	publicKey, err := key.GetPublicKey()
	if err != nil {
		s.deleteUserFromSystem(ctx, username)
	}

	publicKeyString := utils.UnsafeString(publicKey)

	user, err := s.userRepo.Create(ctx, u.User, publicKeyString)
	if err != nil {
		s.deleteUserFromSystem(ctx, username)
		return models.User{}, nil, err
	}

	task, err := tasks.NewUserNotification(u.User, publicKeyString)
	if err != nil {
		s.deleteUserFromDb(ctx, user)
		s.deleteUserFromSystem(ctx, user.Username)
	}

	_, err = s.queue.Enqueue(task)

	if err != nil {
		s.deleteUserFromDb(ctx, user)
		s.deleteUserFromSystem(ctx, user.Username)
	}

	return user, publicKey, err
}

func (s Service) deleteUserFromDb(ctx context.Context, user models.User) {
	err := s.userRepo.Delete(ctx, user.ID)
	s.logger.Error().
		Err(err).
		Uint64("userId", user.ID).
		Msg("Cannot delete newly created user from the database")
}

func (s Service) deleteUserFromSystem(ctx context.Context, userName string) {
	err := s.unixUserService.Delete(ctx, userName)
	s.logger.Error().
		Err(err).
		Str("username", userName).
		Msg("Cannot delete newly created user from the Linux System")
}
