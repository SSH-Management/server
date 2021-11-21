package cli

import (
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/SSH-Management/server/cmd/http/routes"
)

func httpServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start HTTP Server",
		PersistentPreRunE: loadConfig,
		Run:   runHttpServer(),
	}
}

func runHttpServer() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		c := getContainer("logging")

		defer c.Close()

		app := fiber.New(fiber.Config{
			StrictRouting: true,
			AppName:       "SSH Server Management",
		})

		app.Static(
			c.Config.GetString("views.static.path"),
			c.Config.GetString("views.static.dir"),
			fiber.Static{
				Browse:    false,
				Compress:  false,
				ByteRange: true,
			},
		)

		routes.Register(c, app.Group(c.Config.GetString("path_prefix")))

		addr := fmt.Sprintf("%s:%d", c.Config.GetString("bind"), c.Config.GetInt("port"))

		listener, err := net.Listen("tcp4", addr)

		if err != nil {
			log.
				Fatal().
				Err(err).
				Msg("Error while creating net.Listener")
		}

		err = app.Listener(listener)

		if err != nil {
			log.
				Fatal().
				Err(err).
				Msg("Cannot start Fiber HTTP Server")
		}
	}
}
