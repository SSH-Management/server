package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/cmd/http/handlers/auth"
	"github.com/SSH-Management/server/pkg/container"
)

func registerAuthRoutes(c *container.Container, router fiber.Router) {
	router.Post("/login", auth.LoginHandler(
		c.GetLoginService(),
		c.GetValidator(),
		c.GetSession(),
	))
}
