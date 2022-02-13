package routes

import (
	"github.com/SSH-Management/server/pkg/http/handlers/auth"
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/container"
)

func registerAuthRoutes(c *container.Container, router fiber.Router) {
	router.Post("/login", auth.LoginHandler(
		c.GetLoginService(),
		c.GetValidator(),
		c.GetSession(),
	))
}
