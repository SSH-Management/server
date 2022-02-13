package password

import (
	"github.com/stretchr/testify/mock"

	"github.com/SSH-Management/server/pkg/services/password"
)

var _ password.Hasher = &MockHasher{}

type MockHasher struct {
	mock.Mock
}

func (m *MockHasher) Hash(password string) (string, error) {
	args := m.Called(password)

	return args.String(0), args.Error(1)
}

func (m *MockHasher) Verify(hash string, password string) error {
	args := m.Called(hash, password)

	return args.Error(0)
}

func (m *MockHasher) NeedsRehashing(hash string) (bool, error) {
	args := m.Called(hash)

	return args.Bool(0), args.Error(1)
}
