package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Env uint8

const (
	Testing Env = iota
	Development
	Production
)

func ParseEnvironment(env string) Env {
	switch strings.ToLower(env) {
	case "prod", "production":
		return Production
	case "dev", "development":
		return Development
	case "testing", "test":
		return Testing
	default:
		return Production
	}
}

func New(env Env) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName("ssh_management")
	v.SetConfigType("yaml")

	setDefaults(v)

	switch env {
	case Development:
		v.AddConfigPath(".")
	case Production:
		v.AutomaticEnv()
		v.AddConfigPath("/etc/ssh_management/")
	}

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return v, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("database.driver", "mysql")
	v.SetDefault("database.username", "root")
	v.SetDefault("database.password", "password")
	v.SetDefault("database.database", "ssh_management")
	v.SetDefault("database.host", "localhost:3306")
	v.SetDefault("database.collation", "utf8mb4_unicode_ci")
	v.SetDefault("database.conn_max_idle_time", time.Duration(30 * time.Second))
	v.SetDefault("database.conn_max_lifetime", time.Duration(5 * time.Minute))
	v.SetDefault("database.conn_max_idle", 10)
	v.SetDefault("database.conn_max_opened", 10)

	v.SetDefault("bind", "0.0.0.0")
	v.SetDefault("port", 8080)

	v.SetDefault("logging.file", "/var/log/ssh_management/server.log")
	v.SetDefault("logging.level", "info")
	v.SetDefault("logging.console", true)
	v.SetDefault("logging.sample", 0)

	v.SetDefault("crypto.ed25519.private", "/var/keys/private.key")
	v.SetDefault("crypto.ed25519.public", "/var/keys/public.key")

	v.SetDefault("views.static.dir", "./static")
	v.SetDefault("views.static.path", "/static")

	v.SetDefault("redis.host", "localhost")
	v.SetDefault("redis.port", 6379)
	v.SetDefault("redis.username", "")
	v.SetDefault("redis.password", "")

	v.SetDefault("queue.concurrency", 10)
	v.SetDefault("queue.logging.file", "/var/log/ssh_management/queue.log")
	v.SetDefault("queue.logging.level", "info")
	v.SetDefault("queue.logging.console", true)
}
