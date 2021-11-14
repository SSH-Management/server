package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis"
)

func (c *Container) GetStorage(database int) fiber.Storage {
	return redis.New(redis.Config{
		Host:     c.Config.GetString("redis.host"),
		Port:     c.Config.GetInt("redis.port"),
		Username: c.Config.GetString("redis.username"),
		Password: c.Config.GetString("redis.password"),
		Database: database,
	})
}

//
//func (c *Container) GetSession() *session.Store {
//	if c.session == nil {
//		c.session = session.New(session.Config{
//			Storage:        c.GetStorage(0),
//			CookieHTTPOnly: true,
//			//Expiration:     c.Config.Session.Expiration,
//			//KeyLookup:      c.Config.Session.CookieName,
//			//CookieDomain:   c.Config.Session.CookieDomain,
//			//CookiePath:     c.Config.Session.CookiePath,
//			//CookieSecure:   c.Config.Session.Secure,
//			CookieSameSite: "Lax",
//			KeyGenerator: func() string {
//				return utils.RandomString(16)
//			},
//		})
//	}
//
//	return c.session
//}
