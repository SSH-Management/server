package users

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/leebenson/conform"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/http/handlers"
	"github.com/SSH-Management/server/pkg/services/user"
)

func CreateUserHandler(userService user.Interface, validator *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userDto dto.CreateUser

		if err := c.BodyParser(&userDto); err != nil {
			return handlers.ErrInvalidPayload
		}

		if err := conform.Strings(&userDto); err != nil {
			return handlers.ErrInvalidPayload
		}

		if err := validator.Struct(userDto); err != nil {
			return err
		}

		userModel, _, err := userService.Create(c.UserContext(), userDto)

		if err != nil {
			return err
		}

		return c.Status(http.StatusCreated).JSON(userModel)
	}
}
