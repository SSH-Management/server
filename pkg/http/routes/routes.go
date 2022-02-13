package routes

import (
	"net/http"
	"strings"

	"github.com/SSH-Management/server/pkg/config"
	"github.com/SSH-Management/server/pkg/http/handlers"
	"github.com/SSH-Management/server/pkg/http/middleware"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/SSH-Management/server/pkg/container"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// OVDE DEFINISES RUTE
// SANTA MARIA DELLA SALUTE
// ----       STEFAN BOGDANOVIC 2021

func getUiPath(config *config.Config) string {
	uiPath := config.GetString("http.ui.path")

	if !strings.HasSuffix("/", uiPath) {
		uiPath = uiPath + "/"
	}

	return uiPath
}

func Register(c *container.Container, router *fiber.App, environment config.Env) {
	if environment != config.Testing {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     strings.Join(c.Config.GetStringSlice("cors.origins"), ","),
			AllowHeaders:     strings.Join(c.Config.GetStringSlice("cors.headers"), ","),
			AllowMethods:     strings.Join(c.Config.GetStringSlice("cors.methods"), ","),
			AllowCredentials: true,
			MaxAge:           3600,
		}))

		router.Use(middleware.Csrf(
			c.Config,
			c.GetRedisStorage(c.Config.GetInt("redis.csrf.db")),
		))
	}

	router.Get("/csrf-token", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNoContent)
	})

	apiV1 := router.Group("/api/v1")

	registerAuthRoutes(c, apiV1.Group("/auth"))
	registerUserRoutes(c, apiV1.Group("/users"))
	registerServerRoutes(c, apiV1.Group("/servers"))

	if c.Config.GetBool("http.enable_monitor") {
		router.Get("/monitor", monitor.New())
	}

	ui := getUiPath(c.Config)

	router.Static("/", ui, fiber.Static{
		Compress:  true,
		ByteRange: true,
	})

	router.Use(handlers.NotFound(ui))
}
