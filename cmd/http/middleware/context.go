package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

const (
	CancelFuncContextKey         = "cancel"
	CancelWillBeCalledContextKey = "cancelFnWillBeCalled"
)

func Context(ctx *fiber.Ctx) error {
	c, cancel := context.WithCancel(context.Background())

	ctx.Locals(CancelFuncContextKey, cancel)
	ctx.SetUserContext(c)

	err := ctx.Next()

	cancelFnWillBeCalled := ctx.Locals(CancelWillBeCalledContextKey)

	if cancelFnWillBeCalled == nil {
		defer cancel()
	}

	return err
}
