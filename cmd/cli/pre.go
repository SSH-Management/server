package cli

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/SSH-Management/server/pkg/config"
	"github.com/SSH-Management/server/pkg/log"
)

func loadConfig(cmd *cobra.Command, args []string) error {
	log.ConfigureDefaultLogger(LoggingLevel, os.Stdout)

	v, err := config.New(config.ParseEnvironment(Environment))

	if err != nil {
		return err
	}

	viperConfig = v

	return nil
}
