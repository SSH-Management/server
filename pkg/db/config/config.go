package config

import "time"

type Config struct {
	Driver          string        `mapstructure:"driver" json:"driver,omitempty" yaml:"driver,omitempty"`
	Username        string        `mapstructure:"username" json:"username,omitempty" yaml:"username,omitempty"`
	Password        string        `mapstructure:"password" json:"password,omitempty" yaml:"password,omitempty"`
	Database        string        `mapstructure:"database" json:"database,omitempty" yaml:"database,omitempty"`
	Collation       string        `mapstructure:"collation" json:"collation,omitempty" yaml:"collation,omitempty"`
	Host            string        `mapstructure:"host" json:"host,omitempty" yaml:"host,omitempty"`
	MaxIdleTime     time.Duration `mapstructure:"conn_max_idle_time" json:"conn_max_idle_time,omitempty" yaml:"conn_max_idle_time,omitempty"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" json:"conn_max_lifetime,omitempty" yaml:"conn_max_lifetime,omitempty"`
	ConnMaxIdle     int           `mapstructure:"conn_max_idle" json:"conn_max_idle,omitempty" yaml:"conn_max_idle,omitempty"`
	ConnMaxOpen     int           `mapstructure:"conn_max_opened" json:"conn_max_opened,omitempty" yaml:"conn_max_opened,omitempty"`
}
