package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/cmd/http/handlers/client"
	"github.com/SSH-Management/server/cmd/http/middleware"
	"github.com/SSH-Management/server/pkg/container"
)

func registerClientRoutes(c *container.Container, router fiber.Router) {
	router.Use(middleware.Auth(c.GetSession()))

	router.Get("/", client.GetServers(c.GetServerRepository()))
	router.Get("/:id", client.GetServer(c.GetServerRepository()))

	// router.Post("/client/new", handlers.CreateNewClientHandler(
	// 	c.Config.GetString(""),
	// 	c.GetDefaultLogger(),
	// 	c.GetServerRepository(),
	// 	c.GetUserRepository(),
	// ))

	// router.Delete("/client/:id", handlers.DeleteClient(c.GetServerRepository()))
}
