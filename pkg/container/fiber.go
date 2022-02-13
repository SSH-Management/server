package container

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"

	"github.com/SSH-Management/utils/v2"
)

func (c *Container) GetRedisStorage(database int) fiber.Storage {
	return redis.New(redis.Config{
		Host:     c.Config.GetString("redis.host"),
		Port:     c.Config.GetInt("redis.port"),
		Username: c.Config.GetString("redis.username"),
		Password: c.Config.GetString("redis.password"),
		Database: database,
	})
}

func (c *Container) GetSession() *session.Store {
	if c.session == nil {
		c.session = session.New(session.Config{
			Storage:        c.GetRedisStorage(c.Config.GetInt("redis.session.db")),
			CookieHTTPOnly: true,
			Expiration:     c.Config.GetDuration("session.expiration"),
			KeyLookup:      fmt.Sprintf("cookie:%s", c.Config.GetString("session.lookup")),
			CookieDomain:   c.Config.GetString("http.domain"),
			CookiePath:     c.Config.GetString("session.cookie_path"),
			CookieSecure:   c.Config.GetBool("session.secure"),
			CookieSameSite: "strict",
			KeyGenerator: func() string {
				return utils.RandomString(32)
			},
		})
	}

	return c.session
}
