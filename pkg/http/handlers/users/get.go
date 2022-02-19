package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/services/user"
)


func GetUsers(userService user.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := userService.Get(c.UserContext())
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(users)
	}
}

func GetUser(userService user.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	}
}
