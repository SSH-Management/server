package client

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/repositories/server"
)

func GetServers(serverRepo server.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		servers, err := serverRepo.FindAll(c.UserContext())
		if err != nil {
			return err
		}

		return c.JSON(servers)
	}
}

func GetServer(serverRepo server.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return err
		}

		server, err := serverRepo.Find(c.UserContext(), uint64(id))
		if err != nil {
			return err
		}

		return c.JSON(server)
	}
}
