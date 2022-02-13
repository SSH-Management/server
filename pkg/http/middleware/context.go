package middleware

import (
	"context"

	"github.com/SSH-Management/server/pkg/constants"

	"github.com/gofiber/fiber/v2"
)

func Context(ctx *fiber.Ctx) error {
	c, cancel := context.WithCancel(context.Background())

	ctx.Locals(constants.CancelFuncContextKey, cancel)
	ctx.SetUserContext(c)

	err := ctx.Next()

	cancelFnWillBeCalled := ctx.Locals(constants.CancelWillBeCalledContextKey)

	if cancelFnWillBeCalled == nil {
		defer cancel()
	}

	return err
}
