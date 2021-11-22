package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/cmd/http/handlers"
	"github.com/SSH-Management/server/pkg/container"
)


func registerUserRoutes(c *container.Container, router fiber.Router) {
	router.Post("/create", handlers.CreateUserHandler(
		c.GetUserService(),
	))
}
