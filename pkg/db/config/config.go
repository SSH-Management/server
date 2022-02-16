package config

import (
	"fmt"
	"time"
)

type SSLConfig struct{}

type Config struct {
	Username        string        `mapstructure:"username" json:"username,omitempty" yaml:"username,omitempty"`
	Password        string        `mapstructure:"password" json:"password,omitempty" yaml:"password,omitempty"`
	Database        string        `mapstructure:"database" json:"database,omitempty" yaml:"database,omitempty"`
	Host            string        `mapstructure:"host" json:"host,omitempty" yaml:"host,omitempty"`
	Schema          string        `mapstructure:"schema" json:"schema,omitempty" yaml:"schema,omitempty"`
	Port            int           `mapstructure:"port" json:"port,omitempty" yaml:"port,omitempty"`
	TimeZone        string        `mapstructure:"time_zone" json:"time_zone,omitempty" yaml:"time_zone,omitempty"`
	SSLMode         string        `mapstructure:"ssl_mode" json:"ssl_mode,omitempty" yaml:"ssl_mode,omitempty"`
	MaxIdleTime     time.Duration `mapstructure:"conn_max_idle_time" json:"conn_max_idle_time,omitempty" yaml:"conn_max_idle_time,omitempty"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" json:"conn_max_lifetime,omitempty" yaml:"conn_max_lifetime,omitempty"`
	ConnMaxIdle     int           `mapstructure:"conn_max_idle" json:"conn_max_idle,omitempty" yaml:"conn_max_idle,omitempty"`
	ConnMaxOpen     int           `mapstructure:"conn_max_opened" json:"conn_max_opened,omitempty" yaml:"conn_max_opened,omitempty"`
}

func (c Config) FormatConnectionStringURL() string {
	host := "localhost"

	if c.Host != "" {
		host = c.Host
	}

	port := c.Port

	if port == 0 {
		port = 5432
	}

	username := "postgres"

	if c.Username != "" {
		username = c.Username
	}

	password := "postgres"

	if c.Password != "" {
		password = c.Password
	}

	database := "ssh_management"

	if c.Database != "" {
		if c.Database == "(empty)" {
			database = ""
		} else {
			database = c.Database
		}
	}

	schema := "public"

	if c.Schema != "" {
		schema = c.Schema
	}

	sslMode := "disable"

	if c.SSLMode != "" {
		sslMode = c.SSLMode
	}

	timeZone := "UTC"

	if c.TimeZone != "" {
		timeZone = c.TimeZone
	}

	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=%s&search_path=%s",
		username,
		password,
		host,
		port,
		database,
		sslMode,
		timeZone,
		schema,
	)
}

func (c Config) FormatConnectionString() string {
	host := "localhost"

	if c.Host != "" {
		host = c.Host
	}

	port := c.Port

	if port == 0 {
		port = 5432
	}

	username := "postgres"

	if c.Username != "" {
		username = c.Username
	}

	password := "postgres"

	if c.Password != "" {
		password = c.Password
	}

	database := "dbname=ssh_management"

	if c.Database != "" {
		if c.Database == "(empty)" {
			database = ""
		} else {
			database = fmt.Sprintf("dbname=%s", c.Database)
		}
	}

	schema := "public"

	if c.Schema != "" {
		schema = c.Schema
	}

	sslMode := "disable"

	if c.SSLMode != "" {
		sslMode = c.SSLMode
	}

	timeZone := "UTC"

	if c.TimeZone != "" {
		timeZone = c.TimeZone
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s %s port=%d sslmode=%s application_name=%s search_path=%s TimeZone=%s",
		host,
		username,
		password,
		database,
		port,
		sslMode,
		"SSHManagementServer",
		schema,
		timeZone,
	)
}

func (c Config) Clone() Config {
	return Config{
		Username:        c.Username,
		Password:        c.Password,
		Database:        c.Database,
		Host:            c.Host,
		Port:            c.Port,
		TimeZone:        c.TimeZone,
		SSLMode:         c.SSLMode,
		MaxIdleTime:     c.MaxIdleTime,
		ConnMaxLifetime: c.ConnMaxLifetime,
		ConnMaxIdle:     c.ConnMaxIdle,
		ConnMaxOpen:     c.ConnMaxOpen,
	}
}
