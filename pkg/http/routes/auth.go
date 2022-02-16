package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/http/handlers/auth"
	"github.com/SSH-Management/server/pkg/http/middleware"
)

func registerAuthRoutes(c *container.Container, router fiber.Router) {
	router.Post("/login", auth.LoginHandler(
		c.GetLoginService(),
		c.GetValidator(),
		c.GetSession(),
	))

	router.Post("/logout",
		middleware.Auth(c.GetSession()),
		auth.Logout(c.GetSession()),
	)
}
