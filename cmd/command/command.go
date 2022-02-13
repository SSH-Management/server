package command

import (
	"errors"
	"os"

	signer "github.com/SSH-Management/request-signer/v4"
	"github.com/SSH-Management/server/pkg/config"
	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/log"
	zerologlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	viperConfig *config.Config

	Environment  string
	LoggingLevel string
)

func GetContainer(logger ...string) *container.Container {
	l := "logging"

	if len(logger) > 0 {
		l = logger[0]
	}

	c := container.New(l, viperConfig)

	// Generate Key Pair
	if err := c.GetKeyGenerator().Generate(); err != nil && !errors.Is(err, signer.ErrKeysAlreadyExist) {
		zerologlog.
			Fatal().
			Err(err).
			Msg("Error while generating ed25519 key pair")
	}

	return c
}

func LoadConfig(*cobra.Command, []string) error {
	log.ConfigureDefaultLogger(LoggingLevel, os.Stdout)

	v, err := config.New(config.ParseEnvironment(Environment))
	if err != nil {
		return err
	}

	viperConfig = v

	return nil
}
