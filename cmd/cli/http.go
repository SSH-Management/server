package cli

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/SSH-Management/server/cmd/http/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	zerologlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/SSH-Management/server/cmd/command"
	"github.com/SSH-Management/server/cmd/http/routes"
	"github.com/SSH-Management/server/pkg/container"
	services "github.com/SSH-Management/server/pkg/services/grpc"
	"github.com/SSH-Management/server/pkg/services/grpc/middleware"
)

var signals = [4]os.Signal{
	syscall.SIGTERM,
	syscall.SIGINT,
	syscall.SIGUSR1,
	syscall.SIGUSR2,
}

func httpServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "serve",
		Short:             "Start HTTP and GRPC Server",
		PersistentPreRunE: command.LoadConfig,
		RunE:              runHttpServer(),
	}
}

func runHttpServer() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c := command.GetContainer()
		defer c.Close()

		done := make(chan os.Signal, len(signals))
		defer close(done)

		signal.Notify(done, signals[:]...)

		staticConfig := fiber.Config{
			StrictRouting: true,
			AppName:       "SSH Server Management",
			ErrorHandler: handlers.Error(
				c.GetDefaultLogger().Logger,
				c.GetTranslator(),
			),
		}

		app := fiber.New(staticConfig)

		app.Static(
			c.Config.GetString("views.static.path"),
			c.Config.GetString("views.static.dir"),
			fiber.Static{
				Browse:    false,
				Compress:  false,
				ByteRange: true,
			},
		)

		routes.Register(c, app.Group(c.Config.GetString("http.path_prefix")))

		go runFiberHTTPServer(c, app)

		grpcServer := grpc.NewServer(middleware.Register(c)...)

		services.RegisterServices(c, grpcServer)

		reflection.Register(grpcServer)

		go runGRPCServer(c, grpcServer)

		status := <-done

		grpcServer.GracefulStop()

		if err := app.Shutdown(); err != nil {
			log.Error().
				Msg("Error while stopping GoFiber HTTP Server")

			return err
		}

		log.Info().Msgf("Exiting the application: %d", status)
		return nil
	}
}

func runFiberHTTPServer(c *container.Container, app *fiber.App) {
	addr := fmt.Sprintf("%s:%d",
		c.Config.GetString("http.bind"),
		c.Config.GetInt("http.port"),
	)

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

func runGRPCServer(c *container.Container, grpcServer *grpc.Server) {
	addr := fmt.Sprintf("%s:%d",
		c.Config.GetString("http.grpc.bind"),
		c.Config.GetInt("http.grpc.port"),
	)

	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Error while creating net.Listener for GRPC")
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		zerologlog.Fatal().
			Err(err).
			Msg("error while starting grpc server")
	}
}
