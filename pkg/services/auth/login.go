package auth

import (
	"context"

	"github.com/SSH-Management/server/pkg/models"
	"github.com/SSH-Management/server/pkg/repositories/user"
	"github.com/SSH-Management/server/pkg/services/password"
)

type LoginService struct {
	userRepo user.Interface
	hasher   password.Hasher
}

func NewLoginService(userRepo user.Interface, hasher password.Hasher) *LoginService {
	return &LoginService{
		userRepo: userRepo,
		hasher:   hasher,
	}
}

func (s LoginService) Login(ctx context.Context, email, password string) (models.User, error) {
	u, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	if err := s.hasher.Verify(u.Password, password); err != nil {
		return models.User{}, err
	}

	return u, nil
}
