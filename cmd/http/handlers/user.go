package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/leebenson/conform"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/services/user"
)

func CreateUserHandler(userService user.Interface, validator *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userDto dto.CreateUser

		if err := c.BodyParser(&userDto); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Message: "Paylaod is not valid",
			})
		}

		if err := conform.Strings(&userDto); err != nil {
			return c.Status(http.StatusUnprocessableEntity).
				JSON(ErrorResponse{Message: "Payload is not valid"})
		}

		if err := validator.Struct(userDto); err != nil {
			return err
		}

		_, _, err := userService.Create(c.UserContext(), userDto)
		if err != nil {
			return err
		}

		return c.Status(http.StatusCreated).
			JSON(message{Message: "User created"})
	}
}

func GetUsers(userService user.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := userService.Get(c.UserContext())
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(users)
	}
}

func DeleteUser(userService user.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNoContent)
	}
}
