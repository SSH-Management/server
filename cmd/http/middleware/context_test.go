package middleware

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

func TestContextMiddleware(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app := fiber.New()

	app.Use(Context)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	h := app.Handler()

	ctx := &fasthttp.RequestCtx{}

	h(ctx)

	assert.NotNil(ctx.UserValue(CancelFuncContextKey))
}
