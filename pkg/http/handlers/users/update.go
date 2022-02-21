package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/services/user"
)


func UpdateUser(userService user.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	}
}
