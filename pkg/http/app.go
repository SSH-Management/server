package http

import (
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog/log"

	"github.com/SSH-Management/utils/v2"

	"github.com/SSH-Management/server/pkg/config"
	"github.com/SSH-Management/server/pkg/constants"
	"github.com/SSH-Management/server/pkg/http/middleware"
)

func CreateApplication(viewsPath, viewsDir string, static bool, environment config.Env, errorHandler fiber.ErrorHandler) *fiber.App {
	staticConfig := fiber.Config{
		StrictRouting: true,
		AppName:       "SSH Server Management",
		ErrorHandler:  errorHandler,
	}

	app := fiber.New(staticConfig)

	switch environment {
	case config.Development:
		app.Use(pprof.New())
	case config.Production:
		app.Use(recover.New())
	}
	if static {
		app.Static(
			viewsPath,
			viewsDir,
			fiber.Static{
				Browse:    false,
				Compress:  true,
				ByteRange: true,
				MaxAge:    7200,
				Download:  false,
			},
		)

	}

	app.Use(middleware.Context)
	app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return utils.RandomString(32)
		},
		ContextKey: constants.RequestIdKey,
	}))

	return app
}

func RunServer(ip string, port int, app *fiber.App) {
	addr := fmt.Sprintf("%s:%d", ip, port)

	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Error while creating net.Listener for HTTP Server")
	}

	err = app.Listener(listener)

	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Cannot start Fiber HTTP Server")
	}
}
