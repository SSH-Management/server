package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/http/handlers/groups"
)

func registerGroupRoutes(c *container.Container, router fiber.Router) {
	router.Get("/", groups.GetGroups(c.GetGroupRepository()))
}
