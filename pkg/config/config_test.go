package config

import (
	"github.com/spf13/viper"
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

// TODO: Test public key loading
//t.Run("PublicKeyError", func(t *testing.T) {
//	c, err := New(
//		Testing,
//		WithPath("./testdata"),
//		WithConfigFileName("ssh_management_error"),
//		WithConfigType("yaml"),
//	)
//
//	assert.Error(err)
//	_, ok := err.(*os.PathError)
//	assert.True(ok)
//	assert.Nil(c)
//})

func TestNew(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("Ok", func(t *testing.T) {
		c, err := New(
			Testing,
			WithPath("./testdata"),
			WithConfigFileName("ssh_management_ok"),
			WithConfigType("yaml"),
		)

		assert.NoError(err)
		assert.NotNil(c)
		assert.Equal("development", c.GetString("environment"))
	})

	t.Run("NoConfigFile", func(t *testing.T) {
		c, err := New(
			Testing,
			WithPath("./testdata"),
			WithConfigFileName("ssh_management_not_found"),
			WithConfigType("yaml"),
		)

		assert.Error(err)
		_, ok := err.(viper.ConfigFileNotFoundError)
		assert.True(ok)
		assert.Nil(c)
	})

	t.Run("InvalidConfig", func(t *testing.T) {
		c, err := New(
			Testing,
			WithPath("./testdata"),
			WithConfigFileName("ssh_management_invalid"),
			WithConfigType("yaml"),
		)

		assert.Error(err)
		_, ok := err.(viper.ConfigParseError)
		assert.True(ok)
		assert.Nil(c)
	})
}
