package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Logout(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		if err := sess.Destroy(); err != nil {
			return err
		}

		defer func() {
			_ = sess.Save()
		}()

		return c.SendStatus(fiber.StatusNoContent)
	}
}
