package log

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	type Data struct {
		Value    string
		Expected zerolog.Level
	}

	data := []Data{
		{
			Value:    "panic",
			Expected: zerolog.PanicLevel,
		},
		{
			Value:    "fatal",
			Expected: zerolog.FatalLevel,
		},
		{
			Value:    "error",
			Expected: zerolog.ErrorLevel,
		},
		{
			Value:    "warn",
			Expected: zerolog.WarnLevel,
		},
		{
			Value:    "debug",
			Expected: zerolog.DebugLevel,
		},
		{
			Value:    "trace",
			Expected: zerolog.TraceLevel,
		},
		{
			Value:    "info",
			Expected: zerolog.InfoLevel,
		},
		{
			Value:    "other",
			Expected: zerolog.Disabled,
		},
		{
			Value:    "TrAcE",
			Expected: zerolog.TraceLevel,
		},
	}

	for _, item := range data {
		assert.Equal(item.Expected, Parse(item.Value))
	}
}
