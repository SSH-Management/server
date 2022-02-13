package auth

import (
	"time"

	"github.com/SSH-Management/server/pkg/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/leebenson/conform"
	zerologlog "github.com/rs/zerolog/log"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/services/auth"
)

func LoginHandler(loginService *auth.LoginService, validator *validator.Validate, store *session.Store) fiber.Handler {
	store.RegisterType(models.User{})

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

		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		if !sess.Fresh() {
			if err = sess.Regenerate(); err != nil {
				return err
			}
		}

		user, err := loginService.Login(c.UserContext(), login.Email, login.Password)
		if err != nil {
			_ = sess.Destroy()
			return err
		}

		defer func() {
			err = sess.Save()
			zerologlog.Error().Err(err).Msg("Failed to save")
		}()

		sess.Set("user", user)
		sess.SetExpiry(time.Hour)

		return c.JSON(user)
	}
}
