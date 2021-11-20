package cli

import (
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/SSH-Management/server/cmd/http/routes"
	"github.com/SSH-Management/server/pkg/container"
)

func httpServerCommand(c *container.Container, v *viper.Viper) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start HTTP Server",
		Run:   runHttpServer(c, v),
	}
}

func runHttpServer(c *container.Container, v *viper.Viper) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		app := fiber.New(fiber.Config{
			StrictRouting: true,
			AppName:       "SSH Server Management",
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
