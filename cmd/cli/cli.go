package cli

import (
	"context"
	"embed"

	zerologlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/SSH-Management/server/cmd/cli/user"
	"github.com/SSH-Management/server/cmd/command"
)

var rootCmd *cobra.Command

func Execute(version string, migrations embed.FS) {
	rootCmd = &cobra.Command{
		Use:               "server",
		Short:             "SSH Server",
		Long:              `SSH Server Manager - Manages users on instances across clouds`,
		PersistentPreRunE: command.LoadConfig,
		Version:           version,
	}

	flags := rootCmd.PersistentFlags()

	flags.StringVarP(&command.LoggingLevel, "logging-level", "l", "info", "Global Logging level")
	flags.StringVarP(&command.Environment, "env", "e", "production", "Running Environment (Production|Development|Testing)")

	rootCmd.AddCommand(httpServerCommand())
	rootCmd.AddCommand(queueWorkerCommand())
	rootCmd.AddCommand(user.UserCommand())
	rootCmd.AddCommand(migrateCommand(migrations))

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		zerologlog.Fatal().
			Err(err).
			Msg("Error while running command")
	}
}
