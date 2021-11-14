package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/cmd/http/handlers"
	"github.com/SSH-Management/server/pkg/container"
)

func registerClientRoutes(c *container.Container, router fiber.Router) {
	router.Post("/client/new", handlers.CreateNewClientHandler(
		c.Config.GetString("crypto.ed25519.public"),
		c.Logger,
		c.GetServerRepository(),
		c.GetUserRepository(),
	))

	router.Delete("/client/:id", handlers.DeleteClient(c.GetServerRepository()))
}
