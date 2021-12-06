package cli

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/SSH-Management/server/cmd/command"
	"github.com/SSH-Management/server/cmd/http/handlers"
	"github.com/SSH-Management/server/cmd/http/routes"
	"github.com/SSH-Management/server/pkg/container"
)

func httpServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "serve",
		Short:             "Start HTTP Server",
		PersistentPreRunE: command.LoadConfig,
		RunE:              runHttpServer,
	}
}

func runHttpServer(cmd *cobra.Command, args []string) error {
	c := command.GetContainer()

	defer c.Close()

	done := make(chan os.Signal, 4)
	defer close(done)

	signal.Notify(done, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2)

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		AppName:       "SSH Server Management",
		ErrorHandler: handlers.Error(
			c.GetDefaultLogger().Logger,
			c.GetTranslator(),
		),
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

	go runFiberHTTPServer(c, app)

	status := <-done

	if err := app.Shutdown(); err != nil {
		log.Error().
			Msg("Error while stopping GoFiber HTTP Server")
		return err
	}

	log.Info().Msgf("Exiting the application: %d", status)
	return nil
}

func runFiberHTTPServer(c *container.Container, app *fiber.App) {
	addr := fmt.Sprintf("%s:%d", c.Config.GetString("http.bind"), c.Config.GetInt("http.port"))

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
