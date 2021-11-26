package log

import (
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/SSH-Management/utils/v2"

	"github.com/rs/zerolog"
	zerologlog "github.com/rs/zerolog/log"
	"github.com/rzajac/zltest"
)

const DateTimeFormat = "2006-01-02 15:04:05"

type (
	Logger struct {
		zerolog.Logger
		Level zerolog.Level

		FilePath string
		File     *os.File
	}

	UnixServiceLogger struct {
		Logger *Logger
	}
)

func Parse(level string) zerolog.Level {
	level = strings.ToLower(level)

	switch level {
	case "panic":
		return zerolog.PanicLevel
	case "fatal":
		return zerolog.FatalLevel
	case "error":
		return zerolog.ErrorLevel
	case "warn":
		return zerolog.WarnLevel
	case "debug":
		return zerolog.DebugLevel
	case "trace":
		return zerolog.TraceLevel
	case "info":
		return zerolog.InfoLevel
	}

	return zerolog.Disabled
}

func (l Logger) Close() error {
	return l.File.Close()
}

func (l UnixServiceLogger) Print(msg string, data ...interface{}) {
	l.Logger.
		Error().
		Msgf(msg, data...)
}

func ConfigureDefaultLogger(level string, writer io.Writer) {
	zerolog.SetGlobalLevel(Parse(level))
	zerolog.TimeFieldFormat = DateTimeFormat
	zerolog.DurationFieldUnit = time.Microsecond
	zerolog.TimestampFunc = time.Now().UTC

	zerologlog.Logger = zerologlog.Output(zerolog.NewConsoleWriter())
}

func New(fp, level string, toConsole bool, sample uint32) (*Logger, error) {
	file, err := utils.CreateLogFile(fp)
	if err != nil {
		zerologlog.Error().Err(err).Msgf("Error while creating %s log", fp)
		return nil, err
	}

	var logger zerolog.Logger

	if toConsole {
		writers := [2]io.Writer{
			zerolog.NewConsoleWriter(),
			file,
		}

		logger = zerolog.New(zerolog.MultiLevelWriter(writers[:]...)).
			With().
			Timestamp().
			Logger().
			Level(Parse(level))

	} else {
		logger = zerolog.New(file).
			With().
			Timestamp().
			Logger().
			// Sample(&zerolog.BasicSampler{N: sample}).
			Level(Parse(level))
	}

	return &Logger{
		FilePath: fp,
		File:     file,
		Logger:   logger,
		Level:    Parse(level),
	}, err
}

func NewTest(t *testing.T, level zerolog.Level) (*Logger, *zltest.Tester) {
	tst := zltest.New(t)

	logger := zerolog.New(tst).Sample(&zerolog.BasicSampler{N: 1}).Level(level)

	return &Logger{
		Level:  level,
		File:   nil,
		Logger: logger,
	}, tst
}
