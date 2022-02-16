package auth_test

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/SSH-Management/server/pkg/constants"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/SSH-Management/server/pkg/__mocks__/repositories/user"
	"github.com/SSH-Management/server/pkg/__mocks__/services/password"
	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/helpers"
	"github.com/SSH-Management/server/pkg/http/handlers/auth"
	"github.com/SSH-Management/server/pkg/models"
	authservice "github.com/SSH-Management/server/pkg/services/auth"
	passwordservice "github.com/SSH-Management/server/pkg/services/password"
)

func TestLogin_InvalidPassword(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	validator, _ := helpers.GetValidator()

	userRepoMock := new(user.MockUserRepo)
	hasherMock := new(password.MockHasher)

	app, store := helpers.CreateApplication()

	app.Post("/login",
		auth.LoginHandler(
			authservice.NewLoginService(userRepoMock, hasherMock),
			validator,
			store,
		),
	)

	loginDto := dto.Login{
		Email:    "test@test.com",
		Password: "test_password123",
	}

	userRepoMock.On("FindByEmail", mock.Anything, loginDto.Email).
		Once().
		Return(models.User{
			Password: "password_hash",
		}, nil)

	hasherMock.On("Verify", "password_hash", loginDto.Password).
		Once().
		Return(passwordservice.ErrPasswordMismatch)

	res := helpers.Post(app, "/login", helpers.WithBody(loginDto))

	assert.Equal(http.StatusInternalServerError, res.StatusCode)

	assert.Len(res.Cookies(), 1)
	cookie := res.Cookies()[0]
	assert.Equal(helpers.SessionCookieName, cookie.Name)
	assert.Equal("", cookie.Value)
	// Assert Cookie is destroyed
	assert.Less(cookie.Expires.Unix(), time.Now().Unix())
	userRepoMock.AssertExpectations(t)
	hasherMock.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	validator, _ := helpers.GetValidator()

	app, store := helpers.CreateApplication()

	userRepoMock := new(user.MockUserRepo)
	hasherMock := new(password.MockHasher)

	app.Post("/login",
		auth.LoginHandler(
			authservice.NewLoginService(userRepoMock, hasherMock),
			validator,
			store,
		),
	)

	loginDto := dto.Login{
		Email:    "test@test.com",
		Password: "test_password123",
	}

	userRepoMock.On("FindByEmail", mock.Anything, loginDto.Email).
		Once().
		Return(models.User{
			Name:     "Test",
			Surname:  "Test",
			Password: "password_hash",
			Role:     models.NewRole("Administrator", []string{"see:users"}),
		}, nil)

	hasherMock.On("Verify", "password_hash", loginDto.Password).
		Once().
		Return(nil)

	res := helpers.Post(app, "/login", helpers.WithBody(loginDto))

	assert.Equal(http.StatusOK, res.StatusCode)
	assert.Equal("application/json", res.Header.Get("Content-Type"))

	assert.Len(res.Cookies(), 1)

	cookie := res.Cookies()[0]
	assert.Equal(helpers.SessionCookieName, cookie.Name)
	assert.Regexp(regexp.MustCompile(`^[a-zA-Z\d_-]+$`), cookie.Value)

	var resBody auth.LoginResponse

	assert.NoError(json.NewDecoder(res.Body).Decode(&resBody))

	assert.Equal("", resBody.User.Password)
	assert.Equal("Test", resBody.User.Name)
	assert.Equal("Test", resBody.User.Surname)
	assert.Equal("Administrator", resBody.Role)
	assert.EqualValues([]string{"see:users"}, resBody.Permissions)

	sess, _ := helpers.GetSessionWithNameAndValue(app, store, cookie.Name, cookie.Value)

	assert.IsType(models.User{}, sess.Get(constants.SessionUserKey))
	assert.IsType([]string{}, sess.Get(constants.SessionUserPermissionsKey))

	assert.EqualValues(
		[]string{"see:users"},
		sess.Get(constants.SessionUserPermissionsKey).([]string),
	)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	userRepoMock.AssertExpectations(t)
	hasherMock.AssertExpectations(t)
}
