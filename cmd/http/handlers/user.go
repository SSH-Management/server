package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/services/user"
)

func CreateUserHandler(userService user.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userDto dto.CreateUser

		if err := c.BodyParser(&userDto); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Message: "Paylaod is not valid",
			})
		}

		// TODO: Add validation

		_, _, err := userService.Create(c.UserContext(), userDto)

		if err != nil {
			return err
		}

		return c.Status(http.StatusCreated).JSON(message{Message: "User created"})
	}
}
