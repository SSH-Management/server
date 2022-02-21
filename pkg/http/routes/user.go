package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/http/handlers/users"
	"github.com/SSH-Management/server/pkg/http/middleware"
)

func registerUserRoutes(c *container.Container, router fiber.Router) {
	router.Use(middleware.Auth(c.GetSession()))

	router.Get("/", users.GetUsers(
		c.GetUserService(),
	))

	router.Get("/:id", users.GetUser(
		c.GetUserService(),
	))

	router.Post("/create", users.CreateUserHandler(
		c.GetUserService(),
		c.GetValidator(),
	))

	router.Delete("/:id", users.UpdateUser(c.GetUserService()))

	router.Delete("/:id", users.DeleteUser(c.GetUserService()))
}
