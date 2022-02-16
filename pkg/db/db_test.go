package db

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/SSH-Management/server/pkg/db/config"
)

func TestFormatConnectionString(t *testing.T) {
	t.Parallel()
}

func TestGetDbConnection(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")

	if username == "" {
		username = "postgres"
	}

	if password == "" {
		password = "postgres"
	}

	if host == "" {
		host = "localhost"
	}

	var port int64 = 5432

	if portStr != "" {
		port, _ = strconv.ParseInt(portStr, 10, 32)
	}

	db, err := GetDbConnection(config.Config{
		Username:        username,
		Password:        password,
		Database:        "ssh_management",
		Port:            int(port),
		Host:            host,
		MaxIdleTime:     time.Second,
		ConnMaxLifetime: time.Second,
		ConnMaxIdle:     1,
		ConnMaxOpen:     10,
		TimeZone:        "UTC",
	})

	assert.NoError(err)
	assert.NotNil(db)
}
