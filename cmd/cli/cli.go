package cli

import (
	"context"

	zerologlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/SSH-Management/server/cmd/cli/user"
	"github.com/SSH-Management/server/cmd/command"
)

var (
	rootCmd *cobra.Command
)

func Execute() {
	rootCmd = &cobra.Command{
		Use:               "server",
		Short:             "SSH Server",
		Long:              `SSH Server Manager - Manages users on instances across clouds`,
		PersistentPreRunE: command.LoadConfig,
	}

	flags := rootCmd.PersistentFlags()

	flags.StringVarP(&command.Environment, "logging-level", "l", "info", "Global Logging level")
	flags.StringVarP(&command.Environment, "env", "e", "production", "Running Environment (Production|Development|Testing")

	rootCmd.AddCommand(httpServerCommand())
	rootCmd.AddCommand(queueWorkerCommand())
	rootCmd.AddCommand(user.UserCommand())

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		zerologlog.Fatal().
			Err(err).
			Msg("Error while running command")
	}
}
