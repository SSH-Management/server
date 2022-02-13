package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SSH-Management/server/pkg/helpers"

	"github.com/SSH-Management/server/pkg/http/handlers"
	"github.com/SSH-Management/server/pkg/services/password"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"github.com/SSH-Management/server/pkg/db"
)

func setupErrorHandlerApp() (*fiber.App, *validator.Validate) {
	v, translations := helpers.GetValidator()

	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.Error(log.Logger, translations),
	})

	return app, v
}

func TestErrorHandler_ReturnFiberError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _ := setupErrorHandlerApp()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return fiber.ErrBadGateway
	})

	m := struct {
		Message string `json:"message"`
	}{}
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))

	assert.Nil(err)
	assert.EqualValues(fiber.StatusBadGateway, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
	assert.Nil(json.NewDecoder(res.Body).Decode(&m))
	assert.NotEmpty(m.Message)
}

func TestErrorHandler_ValidationError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app, _ := setupErrorHandlerApp()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return validator.ValidationErrors{}
	})
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Nil(err)
	assert.EqualValues(fiber.StatusUnprocessableEntity, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
}

func TestErrorHandler_InvalidValidationError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _ := setupErrorHandlerApp()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return &validator.InvalidValidationError{}
	})
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Nil(err)
	assert.EqualValues(fiber.StatusUnprocessableEntity, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
}

func TestErrorHandler(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app, _ := setupErrorHandlerApp()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return db.ErrNotFound
	})
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Nil(err)
	assert.EqualValues(fiber.StatusNotFound, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
}

func TestErrorHandler_AnyError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app, _ := setupErrorHandlerApp()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return errors.New("any other error")
	})
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Nil(err)
	assert.EqualValues(fiber.StatusInternalServerError, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
}

func TestErrorHandler_PasswordError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app, _ := setupErrorHandlerApp()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return password.ErrPasswordMismatch
	})
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.Nil(err)
	assert.EqualValues(fiber.StatusUnauthorized, res.StatusCode)
	assert.EqualValues(fiber.MIMEApplicationJSON, res.Header.Get(fiber.HeaderContentType))
}
