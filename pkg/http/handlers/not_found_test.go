package handlers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/SSH-Management/server/pkg/helpers"
	"github.com/stretchr/testify/require"

	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/http/handlers"
)

func setupNotFoundApplication() *fiber.App {
	app := fiber.New()

	app.Use(handlers.NotFound("./static"))

	return app
}

func TestNotFound_JsonResponse(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app := setupNotFoundApplication()

	res := helpers.Get(app, "/", helpers.WithHeaders(http.Header{
		"Accept": []string{"application/json"},
	}))

	assert.Equal(http.StatusNotFound, res.StatusCode)
	assert.Equal("application/json", res.Header.Get("Content-Type"))

	var errRes handlers.ErrorResponse
	_ = json.NewDecoder(res.Body).Decode(&errRes)

	assert.EqualValues(handlers.ErrorResponse{
		Message: "Page is not found",
	}, errRes)
}
