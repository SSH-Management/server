package routes

import (
	"github.com/SSH-Management/server/pkg/http/handlers"
	"github.com/SSH-Management/server/pkg/http/middleware"
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/container"
)

func registerUserRoutes(c *container.Container, router fiber.Router) {
	router.Use(middleware.Auth(c.GetSession()))

	router.Get("/", handlers.GetUsers(
		c.GetUserService(),
	))

	router.Get("/:id", handlers.GetUser(
		c.GetUserService(),
	))

	router.Post("/create", handlers.CreateUserHandler(
		c.GetUserService(),
		c.GetValidator(),
	))

	router.Delete("/:id", handlers.UpdateUser(c.GetUserService()))

	router.Delete("/:id", handlers.DeleteUser(c.GetUserService()))
}
