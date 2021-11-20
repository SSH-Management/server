package main

import (
	"fmt"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	zerologlog "github.com/rs/zerolog/log"

	"github.com/SSH-Management/server/cmd/http/routes"
	"github.com/SSH-Management/server/pkg/config"
	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/log"
)

var Version = "dev"

func main() {
	log.ConfigureDefaultLogger("info", os.Stdout)

	v, err := config.New(config.Development)

	if err != nil {
		zerologlog.Fatal().Err(err).Msg("Error while parsing config")
	}

	c, err := container.New(v)

	if err != nil {
		zerologlog.
			Fatal().
			Err(err).
			Msg("Error while creating DI Container")
	}

	// Generate Key Pair
	if err := c.GetKeyGenerator().Generate(); err != nil {
		zerologlog.
			Fatal().
			Err(err).
			Msg("Error while generating ed25519 key pair")
	}

	defer c.Close()

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		AppName:       "Mokey V2",
		ViewsLayout:   "layout",
	})

	app.Static(
		v.GetString("views.static.path"),
		v.GetString("views.static.dir"),
		fiber.Static{
			Browse:    false,
			Compress:  false,
			ByteRange: true,
		},
	)

	routes.Register(c, app.Group(v.GetString("path_prefix")))

	addr := fmt.Sprintf("%s:%d", v.GetString("bind"), v.GetInt("port"))

	listener, err := net.Listen("tcp4", addr)

	if err != nil {
		zerologlog.
			Fatal().
			Err(err).
			Msg("Error while creating net.Listener")
	}

	err = app.Listener(listener)

	if err != nil {
		zerologlog.
			Fatal().
			Err(err).
			Msg("Cannot start Fiber HTTP Server")
	}
}
