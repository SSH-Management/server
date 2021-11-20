package main

import (
	"errors"
	"os"

	zerologlog "github.com/rs/zerolog/log"

	signer "github.com/SSH-Management/request-signer/v3"
	"github.com/SSH-Management/server/cmd/cli"
	"github.com/SSH-Management/server/pkg/config"
	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/log"
)

var Version = "dev"

func main() {
	log.ConfigureDefaultLogger("info", os.Stdout)

	v, err := config.New(config.Development)

	if err != nil {
		zerologlog.Fatal().Err(err).Msg("Error while parsing config")
	}

	c, err := container.New(v)

	if err != nil {
		zerologlog.
			Fatal().
			Err(err).
			Msg("Error while creating DI Container")
	}

	// Generate Key Pair
	if err := c.GetKeyGenerator().Generate(); err != nil && !errors.Is(err, signer.ErrKeysAlreadyExist) {
		zerologlog.
			Fatal().
			Err(err).
			Msg("Error while generating ed25519 key pair")
	}

	defer c.Close()

	cli.Execute(c, v)
}
