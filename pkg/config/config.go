package config

import (
	"encoding/base64"
	"io/fs"
	"os"
	"strings"

	"github.com/SSH-Management/server/pkg/constants"

	"github.com/spf13/viper"
)

type (
	Env uint8

	Config struct {
		*viper.Viper
		PublicKey string
		Env       Env
	}

	Modifier func(*viper.Viper, Env) *viper.Viper
)

const (
	Testing Env = iota
	Development
	Production
)

func ParseEnvironment(env string) Env {
	switch strings.ToLower(env) {
	case "prod", "production":
		return Production
	case "dev", "development", "develop":
		return Development
	case "testing", "test":
		return Testing
	default:
		return Production
	}
}

func (c *Config) LoadServerPublicSSHKey() (string, error) {
	if c.PublicKey != "" {
		return c.PublicKey, nil
	}

	keysFs := os.DirFS(c.GetString("crypto.ed25519"))

	contents, err := fs.ReadFile(keysFs, constants.PublicKeyFileName)

	if err != nil {
		return "", err
	}

	c.PublicKey = base64.RawURLEncoding.EncodeToString(contents)

	return c.PublicKey, nil
}

func WithConfigFileName(name string) Modifier {
	return func(v *viper.Viper, env Env) *viper.Viper {
		v.SetConfigName(name)
		return v
	}
}

func WithConfigType(t string) Modifier {
	return func(v *viper.Viper, env Env) *viper.Viper {
		v.SetConfigType(t)
		return v
	}
}

func WithPath(path string) Modifier {
	return func(v *viper.Viper, env Env) *viper.Viper {
		v.AddConfigPath(path)

		return v
	}
}

func WithEnvSupport() Modifier {
	return func(v *viper.Viper, env Env) *viper.Viper {
		v.SetEnvPrefix("SSH_MANAGEMENT")
		v.AutomaticEnv()

		return v
	}
}

func WithDefaultPaths() Modifier {
	return func(v *viper.Viper, env Env) *viper.Viper {
		switch env {
		case Development:
			v.AddConfigPath(".")
		case Production:
			v.AddConfigPath("/etc/ssh_management/")
		}

		return v
	}
}

var defaults = [4]Modifier{
	WithConfigFileName("ssh_management"),
	WithConfigType("yaml"),
	WithDefaultPaths(),
	WithEnvSupport(),
}

func New(env Env, modifiers ...Modifier) (*Config, error) {
	v := viper.New()

	if len(modifiers) == 0 {
		modifiers = defaults[:]
	}

	for _, modifier := range modifiers {
		v = modifier(v, env)
	}

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return &Config{
		Env:       env,
		Viper:     v,
		PublicKey: "",
	}, nil
}
