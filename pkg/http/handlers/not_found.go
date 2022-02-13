package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NotFound(uiPath string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if strings.HasPrefix(ctx.Get("Accept", "application/json"), "application/json") {
			return ctx.
				Status(fiber.StatusNotFound).
				JSON(ErrorResponse{Message: "Page is not found"})
		}

		return ctx.SendFile(uiPath+"index.html", true)
	}
}
