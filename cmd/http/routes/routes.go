package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
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
	RequestIdKey = "request_id"
	CsrfTokenKey = "csrf_token"
)

func CsrfMiddleware(c *viper.Viper, storage fiber.Storage) fiber.Handler {
	return csrf.New(csrf.Config{
		KeyLookup:      fmt.Sprintf("%s:%s", c.GetString("csrf.lookup_method"), c.GetString("csrf.lookup_key")),
		ContextKey:     CsrfTokenKey,
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

	uiPath := c.Config.GetString("http.ui.path")

	if !strings.HasSuffix("/", uiPath) {
		uiPath = uiPath + "/"
	}

	router.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return utils.RandomString(32)
		},
		ContextKey: RequestIdKey,
	}))

	router.Get("/csrf-token", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNoContent)
	})

	apiV1 := router.Group("/api/v1")

	registerClientRoutes(c, apiV1)
	registerAuthRoutes(c, apiV1.Group("/auth"))
	registerUserRoutes(c, apiV1.Group("/users"))

	router.Get("/monitor", monitor.New())

	router.Static("/", uiPath, fiber.Static{
		Compress:  true,
		ByteRange: true,
	})

	router.Use(func(ctx *fiber.Ctx) error {
		if strings.HasPrefix(ctx.Get("Accept", "text/html"), "application/json") {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Route not found",
			})
		}

		return ctx.SendFile(uiPath+"index.html", true)
	})
}
