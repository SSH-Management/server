package group

import (
	"context"

	"github.com/SSH-Management/server/pkg/repositories/group"

	"github.com/stretchr/testify/mock"

	"github.com/SSH-Management/server/pkg/models"
)

var _ group.Interface = &MockGroupRepository{}

type MockGroupRepository struct {
	mock.Mock
}

func (m *MockGroupRepository) Find(ctx context.Context, id uint64) (models.Group, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(models.Group), args.Error(1)
}

func (m *MockGroupRepository) FindByName(ctx context.Context, name ...string) ([]models.Group, error) {
	args := m.Called(ctx, name)

	return args.Get(0).([]models.Group), args.Error(1)
}

func (m *MockGroupRepository) Create(ctx context.Context, name string) (models.Group, error) {
	args := m.Called(ctx, name)

	return args.Get(0).(models.Group), args.Error(1)
}

func (m *MockGroupRepository) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)

	return args.Error(1)
}
