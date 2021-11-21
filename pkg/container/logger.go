package container

import (
	"fmt"

	"github.com/SSH-Management/server/pkg/log"
	zerologlog "github.com/rs/zerolog/log"
)

type LoggerConfig struct {
	File      string
	Level     string
	ToConsole bool
}

func (c *Container) GetDefaultLogger() *log.Logger {
	return c.GetDefaultLogger()
}

func (c *Container) GetLogger(name string) *log.Logger {
	if logger, ok := c.loggers[name]; ok {
		return logger
	}

	logger, err := log.New(
		c.Config.GetString(fmt.Sprintf("%s:file", name)),
		c.Config.GetString(fmt.Sprintf("%s:level", name)),
		c.Config.GetBool(fmt.Sprintf("%s:console", name)),
		0,
	)

	if err != nil {
		zerologlog.Fatal().Err(err).Str("name", name).Msg("Error while creating logger")
	}

	c.loggers[name] = logger

	return logger
}
