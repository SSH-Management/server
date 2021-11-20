package cli

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/SSH-Management/server/pkg/container"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "SSH Server",
		Long:  `SSH Server Manager - Manages users on instances across clouds`,
	}
}

func Execute(c *container.Container, v *viper.Viper) {
	rootCmd.AddCommand(httpServerCommand(c, v))

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Error while running command")
	}
}
