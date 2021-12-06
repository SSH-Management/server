package routes

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"

	"github.com/SSH-Management/utils/v2"

	"github.com/SSH-Management/server/cmd/http/middleware"
	"github.com/SSH-Management/server/pkg/container"
)

const (
	RequestIdKey    = "request_id"
	CsrfTokenCookie = "csrf_cookie"
	CsrfTokenKey    = "csrf_token"
)

func CsrfMiddleware(c *viper.Viper, storage fiber.Storage) fiber.Handler {
	return csrf.New(csrf.Config{
		KeyLookup:      "cookie:csrf_cookie",
		ContextKey:     CsrfTokenKey,
		CookieName:     c.GetString("csrf.lookup"),
		Storage:        storage,
		CookieDomain:   c.GetString("http.domain"),
		CookieSecure:   c.GetBool("csrf.secure"),
		Expiration:     c.GetDuration("csrf.expiration"),
		CookieHTTPOnly: true,
		CookiePath:     c.GetString("csrf.cookie_path"),
		CookieSameSite: c.GetString("csrf.same_site"),
		KeyGenerator: func() string {
			return utils.RandomString(32)
		},
	})
}

func Register(c *container.Container, router fiber.Router) {
	router.Use(pprof.New())
	router.Use(recover.New())
	router.Use(middleware.Context)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(c.Config.GetStringSlice("cors.origins"), ","),
		AllowHeaders:     strings.Join(c.Config.GetStringSlice("cors.headers"), ","),
		AllowMethods:     strings.Join(c.Config.GetStringSlice("cors.methods"), ","),
		AllowCredentials: true,
		MaxAge:           3600,
	}))

	router.Use(CsrfMiddleware(
		c.Config,
		c.GetStorage(c.Config.GetInt("redis.csrf.db")),
	))

	router.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return utils.RandomString(32)
		},
		ContextKey: RequestIdKey,
	}))

	router.Get("/", monitor.New())

	router.Get("/csrf-token", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNoContent)
	})

	registerClientRoutes(c, router)
	registerUserRoutes(c, router.Group("/users"))
	registerAuthRoutes(c, router.Group("/auth"))
}
