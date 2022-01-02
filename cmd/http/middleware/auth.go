package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Auth(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		s, err := store.Get(c)
		if err != nil {
			return err
		}

		user := s.Get("user")

		if user != nil {
			return fiber.ErrUnauthorized
		}

		return c.Next()
	}
}
