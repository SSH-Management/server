package unix

import (
	"context"

	"github.com/stretchr/testify/mock"

	linuxuser "github.com/SSH-Management/linux-user"
)

var _ linuxuser.UnixInterface = &MockUnixService{}

type MockUnixService struct {
	mock.Mock
}

func (m *MockUnixService) Find(ctx context.Context, name string) (linuxuser.User, error) {
	args := m.Called(ctx, name)

	if err := args.Error(1); err != nil {
		return linuxuser.User{}, err
	}

	return args.Get(0).(linuxuser.User), nil
}

func (m *MockUnixService) Create(ctx context.Context, creater linuxuser.Creater) (linuxuser.User, error) {
	args := m.Called(ctx, creater)

	if err := args.Error(1); err != nil {
		return linuxuser.User{}, err
	}

	return args.Get(0).(linuxuser.User), nil
}

func (m *MockUnixService) Delete(ctx context.Context, name string) error {
	args := m.Called(ctx, name)

	return args.Error(0)
}
