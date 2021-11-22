package cli

import (
	"errors"

	zerologlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	signer "github.com/SSH-Management/request-signer/v3"

	"github.com/SSH-Management/server/pkg/container"
)

var (
	rootCmd *cobra.Command

	viperConfig *viper.Viper

	Environment string
	LoggingLevel string
)

func getContainer(logger string) *container.Container {
	c := container.New("logging", viperConfig)

	// Generate Key Pair
	if err := c.GetKeyGenerator().Generate(); err != nil && !errors.Is(err, signer.ErrKeysAlreadyExist) {
		zerologlog.
			Fatal().
			Err(err).
			Msg("Error while generating ed25519 key pair")
	}

	return c
}

func Execute() {
	cobra.OnInitialize(func() {

	})

	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "SSH Server",
		Long:  `SSH Server Manager - Manages users on instances across clouds`,
		PersistentPreRunE: loadConfig,
	}

	rootCmd.PersistentFlags().StringVarP(&LoggingLevel, "logging-level", "l", "info", "Global Logging level")
	rootCmd.PersistentFlags().StringVarP(&Environment, "env", "e", "production", "Running Environment (Production|Development|Testing")

	rootCmd.AddCommand(httpServerCommand())
	rootCmd.AddCommand(queueWorkerCommand())

	if err := rootCmd.Execute(); err != nil {
		zerologlog.Fatal().Err(err).Msg("Error while running command")
	}
}
