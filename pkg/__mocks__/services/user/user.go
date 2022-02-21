package user

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/models"
	user_service "github.com/SSH-Management/server/pkg/services/user"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Get(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)

	err := args.Error(1)

	if err != nil {
		return []models.User{}, err
	}

	return args.Get(0).([]models.User), err
}

func (m *MockUserService) Create(ctx context.Context, user dto.CreateUser) (models.User, []byte, error) {
	args := m.Called(ctx, user)

	err := args.Error(2)

	if err != nil {
		return models.User{}, nil, err
	}

	return args.Get(0).(models.User), args.Get(1).([]byte), nil
}

var _ user_service.Interface = &MockUserService{}
