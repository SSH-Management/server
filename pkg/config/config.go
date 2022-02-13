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

	Modifier func(*viper.Viper) *viper.Viper
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

func loadServerPublicSSHKey(keysFs fs.FS) (string, error) {
	contents, err := fs.ReadFile(keysFs, constants.PublicKeyFileName)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(contents), nil
}

func WithConfigFileName(name string) Modifier {
	return func(v *viper.Viper) *viper.Viper {
		v.SetConfigName(name)
		return v
	}
}

func WithConfigType(t string) Modifier {
	return func(v *viper.Viper) *viper.Viper {
		v.SetConfigType(t)
		return v
	}
}

var defaults = [2]Modifier{
	WithConfigFileName("ssh_management"),
	WithConfigType("yaml"),
}

func New(env Env, modifiers ...Modifier) (*Config, error) {
	v := viper.New()

	for _, modifier := range defaults {
		v = modifier(v)
	}

	for _, modifier := range modifiers {
		v = modifier(v)
	}

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

	publicKey, err := loadServerPublicSSHKey(os.DirFS(v.GetString("crypto.ed25519")))
	if err != nil {
		return nil, err
	}

	return &Config{
		Env:       env,
		Viper:     v,
		PublicKey: publicKey,
	}, nil
}
