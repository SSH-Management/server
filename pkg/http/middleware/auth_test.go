package middleware_test

import (
	"net/http"
	"testing"

	"github.com/SSH-Management/server/pkg/constants"
	"github.com/SSH-Management/server/pkg/models"

	"github.com/SSH-Management/server/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestUserIsNotAuthenticated(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _ := helpers.CreateApplicationWithSession()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	req := helpers.Get(app, "/")

	assert.Equal(http.StatusUnauthorized, req.StatusCode)
}

func TestAuthMiddleware_InvalidUserType(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, store := helpers.CreateApplicationWithSession()

	session, clear := helpers.GetSession(app, store)

	defer clear()

	session.Set(constants.SessionUserKey, "user")
	assert.NoError(session.Save())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	req := helpers.Get(app, "/", helpers.WithSessionCookie())

	assert.Equal(http.StatusUnauthorized, req.StatusCode)
}

func TestAuthMiddleware_InvalidUserID(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, store := helpers.CreateApplicationWithSession()

	session, clear := helpers.GetSession(app, store)

	defer clear()

	store.RegisterType(models.User{})
	session.Set(constants.SessionUserKey, models.User{})
	assert.NoError(session.Save())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	req := helpers.Get(app, "/", helpers.WithSessionCookie())

	assert.Equal(http.StatusUnauthorized, req.StatusCode)
}

func TestAuthMiddleware_Success(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, store := helpers.CreateApplicationWithSession()

	session, clear := helpers.GetSession(app, store)

	defer clear()
	store.RegisterType(models.User{})
	session.Set(constants.SessionUserKey, models.User{Model: models.Model{ID: 10}})
	assert.NoError(session.Save())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	req := helpers.Get(app, "/", helpers.WithSessionCookie())

	assert.Equal(http.StatusOK, req.StatusCode)
}
