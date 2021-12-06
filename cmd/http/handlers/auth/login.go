package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/leebenson/conform"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/services/auth"
)

func LoginHandler(loginService *auth.LoginService, validator *validator.Validate, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var login dto.Login

		if err := c.BodyParser(&login); err != nil {
			return err
		}

		if err := conform.Strings(&login); err != nil {
			return err
		}

		if err := validator.Struct(login); err != nil {
			return err
		}

		session, err := store.Get(c)

		if err != nil {
			return err
		}

		if !session.Fresh() {
			if err := session.Regenerate(); err != nil {
				return err
			}
		}

		defer func() {
			_ = session.Save()
		}()

		user, err := loginService.Login(c.UserContext(), login.Email, login.Password)

		if err != nil {
			return err
		}

		session.Set("user", user)

		return c.JSON(user)
	}
}
