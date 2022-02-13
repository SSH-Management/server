package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/SSH-Management/server/pkg/__mocks__/repositories/user"
	"github.com/SSH-Management/server/pkg/__mocks__/services/password"
	"github.com/SSH-Management/server/pkg/models"
)

func TestNewLoginService(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	userRepoMock := new(user.MockUserRepo)
	hasherMock := new(password.MockHasher)

	assert.NotNil(NewLoginService(userRepoMock, hasherMock))
}

func TestLogin_UserNotFound(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	userRepoMock := new(user.MockUserRepo)
	hasherMock := new(password.MockHasher)

	service := LoginService{
		userRepo: userRepoMock,
		hasher:   hasherMock,
	}

	ctx := context.Background()

	userRepoMock.On("FindByEmail", mock.Anything, "test@test.com").
		Once().
		Return(models.User{}, errors.New("user not found"))

	hasherMock.AssertNotCalled(t, "Verify")
	userModel, err := service.Login(ctx, "test@test.com", "password123")

	assert.Error(err)
	assert.EqualError(err, "user not found")
	assert.Empty(userModel)
	userRepoMock.AssertExpectations(t)
	hasherMock.AssertExpectations(t)
}

func TestLogin_InvalidPassword(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	userRepoMock := new(user.MockUserRepo)
	hasherMock := new(password.MockHasher)

	service := LoginService{
		userRepo: userRepoMock,
		hasher:   hasherMock,
	}

	ctx := context.Background()

	userRepoMock.On("FindByEmail", mock.Anything, "test@test.com").
		Once().
		Return(models.User{Password: "test_password_hash"}, nil)

	hasherMock.On("Verify", "test_password_hash", "password123").
		Once().Return(errors.New("invalid password"))

	userModel, err := service.Login(ctx, "test@test.com", "password123")

	assert.Error(err)
	assert.EqualError(err, "invalid password")
	assert.Empty(userModel)
	userRepoMock.AssertExpectations(t)
	hasherMock.AssertExpectations(t)
}

func TestLogin_OK(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	userRepoMock := new(user.MockUserRepo)
	hasherMock := new(password.MockHasher)

	service := LoginService{
		userRepo: userRepoMock,
		hasher:   hasherMock,
	}

	ctx := context.Background()

	userRepoMock.On("FindByEmail", mock.Anything, "test@test.com").
		Once().
		Return(models.User{Email: "test@test.com", Password: "test_password_hash"}, nil)

	hasherMock.On("Verify", "test_password_hash", "password123").
		Once().
		Return(nil)

	userModel, err := service.Login(ctx, "test@test.com", "password123")

	assert.NoError(err)
	assert.NotEmpty(userModel)
	userRepoMock.AssertExpectations(t)
	hasherMock.AssertExpectations(t)
}
