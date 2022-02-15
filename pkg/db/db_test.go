package db

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/SSH-Management/server/pkg/db/config"
)

func TestGetDbConnection_InvalidDriver(t *testing.T) {
	t.Parallel()

	assert := require.New(t)

	db, err := GetDbConnection(config.Config{Driver: "invalid_driver"})

	assert.Error(err)
	assert.Equal("driver 'invalid_driver' is not supported", err.Error())
	assert.Nil(db)
}

func TestGetDbConnection(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	portStr := os.Getenv("MYSQL_PORT")

	if username == "" {
		username = "root"
	}

	if password == "" {
		password = "password"
	}

	if host == "" {
		host = "localhost"
	}

	var port int64 = 3306

	if portStr != "" {
		port, _ = strconv.ParseInt(portStr, 10, 32)
	}

	db, err := GetDbConnection(config.Config{
		Driver:          "mysql",
		Username:        username,
		Password:        password,
		Database:        "ssh_management",
		Collation:       "utf8mb4_unicode_ci",
		Host:            fmt.Sprintf("%s:%d", host, port),
		MaxIdleTime:     time.Second,
		ConnMaxLifetime: time.Second,
		ConnMaxIdle:     1,
		ConnMaxOpen:     10,
	})

	assert.NoError(err)
	assert.NotNil(db)
}
