package mysql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/SSH-Management/server/pkg/db/config"
)

func TestFormatDSN(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("WithDB", func(t *testing.T) {
		dsn := FormatDSN(config.Config{
			Driver:          "mysql",
			Username:        "root",
			Password:        "password",
			Database:        "ssh_management",
			Collation:       "utf8mb4_unicode_ci",
			Host:            "localhost:3306",
			MaxIdleTime:     time.Second,
			ConnMaxLifetime: time.Second,
			ConnMaxIdle:     1,
			ConnMaxOpen:     10,
		}, true)

		assert.Equal("root:password@tcp(localhost:3306)/ssh_management?collation=utf8mb4_unicode_ci&multiStatements=true&parseTime=true", dsn)
	})

	t.Run("WithoutDB", func(t *testing.T) {
		dsn := FormatDSN(config.Config{
			Driver:          "mysql",
			Username:        "root",
			Password:        "password",
			Database:        "ssh_management",
			Collation:       "utf8mb4_unicode_ci",
			Host:            "localhost:3306",
			MaxIdleTime:     time.Second,
			ConnMaxLifetime: time.Second,
			ConnMaxIdle:     1,
			ConnMaxOpen:     10,
		}, false)

		assert.Equal("root:password@tcp(localhost:3306)/?collation=utf8mb4_unicode_ci&multiStatements=true&parseTime=true", dsn)
	})
}

func TestMySql_Connect(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	m := mySql{}

	dialect, err := m.Connect(config.Config{
		Driver:          "mysql",
		Username:        "root",
		Password:        "password",
		Database:        "ssh_management",
		Collation:       "utf8mb4_unicode_ci",
		Host:            "localhost:3306",
		MaxIdleTime:     time.Second,
		ConnMaxLifetime: time.Second,
		ConnMaxIdle:     1,
		ConnMaxOpen:     10,
	})

	assert.NoError(err)
	assert.NotNil(dialect)
}

func TestNew(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	assert.NotNil(New())
}