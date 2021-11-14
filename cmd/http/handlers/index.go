package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Index() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
