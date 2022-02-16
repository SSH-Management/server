package auth_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/SSH-Management/server/pkg/constants"
	"github.com/SSH-Management/server/pkg/helpers"
	"github.com/SSH-Management/server/pkg/http/handlers/auth"
	"github.com/SSH-Management/server/pkg/models"
)

func TestLogout_NoSessionCookie(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, store := helpers.CreateApplicationWithSession()
	app.Post("/", auth.Logout(store))

	res := helpers.Post(app, "/")

	assert.Equal(http.StatusUnauthorized, res.StatusCode)
}

func TestLogout_Success(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, store := helpers.CreateApplicationWithSession()
	app.Post("/", auth.Logout(store))

	store.RegisterType(models.User{})

	session, clear := helpers.GetSession(app, store)

	defer clear()

	session.Set(constants.SessionUserKey, models.User{
		Model: models.Model{
			ID: 1,
		},
	})

	assert.NoError(session.Save())

	res := helpers.Post(app, "/", helpers.WithSessionCookie())

	assert.Equal(http.StatusNoContent, res.StatusCode)
}
