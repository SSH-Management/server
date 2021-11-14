package config

import (
	"time"

	"github.com/spf13/viper"
)

type Env uint8

const (
	Testing Env = iota
	Development
	Production
)

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
	case Testing:
	}

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return v, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("develop", false)

	v.SetDefault("bind", "0.0.0.0")
	v.SetDefault("port", 3000)
	v.SetDefault("path_prefix", "/")

	v.SetDefault("ssl.enabled", false)
	v.SetDefault("ssl.key", "")
	v.SetDefault("ssl.cert", "")
	v.SetDefault("ssl.port", 3001)

	v.SetDefault("password.enabled", false)
	v.SetDefault("password.min_len", 8)

	v.SetDefault("views.templates", "./templates")
	v.SetDefault("views.static.dir", "./static")
	v.SetDefault("views.static.path", "/static")

	v.SetDefault("database.driver", "mysql")
	v.SetDefault("database.username", "mokey")
	v.SetDefault("database.password", "mokey")
	v.SetDefault("database.host", "127.0.0.1:3306")
	v.SetDefault("database.collation", "utf8mb4_unicode_ci")
	v.SetDefault("database.collation", "utf8mb4_unicode_ci")
	v.SetDefault("database.conn_max_idle_time", 30*time.Second)
	v.SetDefault("database.conn_max_lifetime", 5*time.Minute)
	v.SetDefault("database.conn_max_idle", 10)
	v.SetDefault("database.conn_max_opened", 10)

	v.SetDefault("smtp.host", "127.0.0.1")
	v.SetDefault("smtp.port", 25)
	v.SetDefault("smtp.username", "")
	v.SetDefault("smtp.password", "")
	v.SetDefault("smtp.tls", "off")

	v.SetDefault("email.from", "helpdesk@example.edu")
	v.SetDefault("email.signature", "Mr. System Administrator")
	v.SetDefault("email.link_base", "http://localhost:8080") // TODO: Format with Default values from port and bind
	v.SetDefault("email.prefix", "mokey")

	v.SetDefault("pgp.sign", false)
	v.SetDefault("pgp.key", "")
	v.SetDefault("pgp.passphrase", "")

	v.SetDefault("verify.by_admin", false)
	v.SetDefault("verify.by_email", false)

	v.SetDefault("redis.host", "127.0.0.1")
	v.SetDefault("redis.port", 6379)
	v.SetDefault("redis.username", "")
	v.SetDefault("redis.password", "")

	v.SetDefault("logging.file", "./logs/mokey.log")
	v.SetDefault("logging.level", "info")
	v.SetDefault("logging.console", false)
	v.SetDefault("logging.sample", 0)
}
