package cli

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	zerologlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/SSH-Management/server/cmd/command"
	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/http"
	"github.com/SSH-Management/server/pkg/http/handlers"
	"github.com/SSH-Management/server/pkg/http/routes"
	services "github.com/SSH-Management/server/pkg/services/grpc"
	"github.com/SSH-Management/server/pkg/services/grpc/middleware"
)

var signals = [2]os.Signal{
	syscall.SIGTERM,
	syscall.SIGINT,
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

		errorHandler := handlers.Error(
			c.GetDefaultLogger(),
			c.GetTranslator(),
		)

		app := http.CreateApplication(
			c.Config.GetString("views.static.path"),
			c.Config.GetString("views.static.dir"),
			true,
			c.Config.Env,
			errorHandler,
		)

		routes.Register(c, app, c.Config.Env)

		go http.RunServer(c.Config.GetString("http.bind"), c.Config.GetInt("http.port"), app)

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

func runGRPCServer(c *container.Container, grpcServer *grpc.Server) {
	addr := fmt.Sprintf(
		"%s:%d",
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
