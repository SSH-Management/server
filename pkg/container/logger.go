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
	return c.GetLogger(c.defaultLoggerName)
}

func (c *Container) GetLogger(name string) *log.Logger {
	if logger, ok := c.loggers[name]; ok {
		return logger
	}

	file := c.Config.GetString(fmt.Sprintf("%s.file", name))

	logger, err := log.New(
		file,
		c.Config.GetString(fmt.Sprintf("%s.level", name)),
		c.Config.GetBool(fmt.Sprintf("%s.console", name)),
		0,
	)

	if err != nil {
		zerologlog.Fatal().
			Err(err).
			Str("name", name).
			Msg("Error while creating logger")
	}

	c.loggers[name] = logger

	return logger
}
