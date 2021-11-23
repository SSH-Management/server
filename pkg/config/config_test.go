package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseEnvironment(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	type theory struct {
		value    string
		expected Env
	}

	data := []theory{
		{value: "production", expected: Production},
		{value: "prod", expected: Production},
		{value: "Prod", expected: Production},
		{value: "PRODUCTION", expected: Production},
		{value: "Production", expected: Production},
		{value: "", expected: Production},

		{value: "testing", expected: Testing},
		{value: "test", expected: Testing},
		{value: "Test", expected: Testing},
		{value: "TESTING", expected: Testing},
		{value: "Testing", expected: Testing},

		{value: "development", expected: Development},
		{value: "dev", expected: Development},
		{value: "Dev", expected: Development},
		{value: "DEVELOPMENT", expected: Development},
		{value: "Development", expected: Development},
	}

	for _, item := range data {
		assert.Equal(item.expected, ParseEnvironment(item.value))
	}
}
