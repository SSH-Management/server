package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"

	"github.com/SSH-Management/utils/v2"

	"github.com/SSH-Management/server/pkg/config"
	"github.com/SSH-Management/server/pkg/constants"
)

func Csrf(c *config.Config, storage fiber.Storage) fiber.Handler {
	return csrf.New(csrf.Config{
		KeyLookup:      fmt.Sprintf("header:%s", c.GetString("csrf.lookup_key")),
		ContextKey:     constants.CsrfTokenKey,
		CookieName:     c.GetString("csrf.cookie_name"),
		Storage:        storage,
		CookieDomain:   c.GetString("http.domain"),
		CookieSecure:   c.GetBool("csrf.secure"),
		Expiration:     c.GetDuration("csrf.expiration"),
		CookieHTTPOnly: false,
		CookiePath:     c.GetString("csrf.cookie_path"),
		CookieSameSite: "strict",
		KeyGenerator: func() string {
			return utils.RandomString(32)
		},
	})
}
