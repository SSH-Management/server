package user

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/models"
	"github.com/SSH-Management/server/pkg/repositories/user"
)

var _ user.Interface = &MockUserRepo{}

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Find(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)

	users := args.Get(0)

	if users == nil {
		return nil, args.Error(1)
	}

	return users.([]models.User), args.Error(1)
}

func (m *MockUserRepo) FindByGroup(ctx context.Context, id uint64) ([]models.User, error) {
	args := m.Called(ctx, id)

	users := args.Get(0)

	if users == nil {
		return nil, args.Error(1)
	}

	return users.([]models.User), args.Error(1)
}

func (m *MockUserRepo) FindByEmail(ctx context.Context, email string) (models.User, error) {
	args := m.Called(ctx, email)

	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepo) Create(ctx context.Context, user dto.User, s string) (models.User, error) {
	args := m.Called(ctx, user, s)

	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepo) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)

	return args.Error(0)
}
