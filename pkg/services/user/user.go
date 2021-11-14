package user

import (
	"context"

	user "github.com/SSH-Management/linux-user"
	ssh "github.com/SSH-Management/ssh"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
	userepo "github.com/SSH-Management/server/pkg/repositories/user"
)

type (
	Service struct {
		userRepo        userepo.Interface
		unixUserService user.UnixInterface

		logger          *log.Logger
	}

	Interface interface {
		Create(context.Context, dto.User) (models.User, []byte, error)
	}
)

func New(userRepo userepo.Interface, unixUserService user.UnixInterface, logger *log.Logger) Service {
	return Service{
		userRepo:        userRepo,
		unixUserService: unixUserService,
		logger:          logger,
	}
}

func (s Service) Create(ctx context.Context, u dto.User) (models.User, []byte, error) {
	user, err := s.userRepo.Create(ctx, u)

	if err != nil {
		return models.User{}, nil, err
	}

	unixUser, err := s.unixUserService.Create(ctx, u)

	if err != nil {
		s.deleteUserFromDb(ctx, user)
		return models.User{}, nil, err
	}

	key, err := ssh.New(unixUser.UserId, unixUser.GroupId, unixUser.HomeFolder)

	if err != nil {
		s.deleteUserFromDb(ctx, user)
		s.deleteUserFromSystem(ctx, user.Username)
		return models.User{}, nil, err
	}

	if err := key.Write(); err != nil {
		s.deleteUserFromDb(ctx, user)
		s.deleteUserFromSystem(ctx, user.Username)
		return models.User{}, nil, err
	}

	publicKey, err := key.GetPublicKey()

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
