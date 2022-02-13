package server

import (
	"context"

	"github.com/SSH-Management/server/pkg/repositories/server"
	"github.com/stretchr/testify/mock"

	"github.com/SSH-Management/protobuf/server/clients"
	"github.com/SSH-Management/server/pkg/models"
)

var _ server.Interface = &MockServerRepository{}

type MockServerRepository struct {
	mock.Mock
}

func (m *MockServerRepository) FindAll(ctx context.Context) ([]models.Server, error) {
	args := m.Called(ctx)

	return args.Get(0).([]models.Server), args.Error(1)
}

func (m *MockServerRepository) FindByPrivateIP(ctx context.Context, ip string) (models.Server, error) {
	args := m.Called(ctx, ip)

	return args.Get(0).(models.Server), args.Error(1)
}

func (m *MockServerRepository) Find(ctx context.Context, id uint64) (models.Server, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(models.Server), args.Error(1)
}

func (m *MockServerRepository) Create(ctx context.Context, request *clients.CreateClientRequest) (models.Server, error) {
	args := m.Called(ctx, request)

	return args.Get(0).(models.Server), args.Error(1)
}

func (m *MockServerRepository) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)

	return args.Error(0)
}
