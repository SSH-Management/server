package role

import (
	"context"

	"github.com/SSH-Management/server/pkg/repositories/role"
	"github.com/stretchr/testify/mock"

	"github.com/SSH-Management/server/pkg/models"
)

var _ role.Interface = &MockRoleRepository{}

type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) Find(ctx context.Context, id uint64) (models.Role, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(models.Role), args.Error(1)
}

func (m *MockRoleRepository) FindByName(ctx context.Context, name string) (models.Role, error) {
	args := m.Called(ctx, name)

	return args.Get(0).(models.Role), args.Error(1)
}

func (m *MockRoleRepository) Create(ctx context.Context) (models.Role, error) {
	args := m.Called(ctx)

	return args.Get(0).(models.Role), args.Error(1)
}

func (m *MockRoleRepository) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)

	return args.Error(1)
}
