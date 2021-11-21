package container

import (
	"errors"

	signer "github.com/SSH-Management/request-signer/v3"
	"github.com/SSH-Management/utils/v2"
)

func (c *Container) GetSigner() signer.Signer {
	if c.signer == nil {
		var err error

		generator := c.GetKeyGenerator()

		if err := generator.Generate(); err != nil && !errors.Is(err, signer.ErrKeysAlreadyExist) {
			c.GetDefaultLogger().Fatal().
				Err(err).
				Msg("Failed to generate ed25519 keys")
		}

		c.signer, err = signer.NewSigner(c.Config.GetString("crypto.ed25519.private"))

		if err != nil {
			c.GetDefaultLogger().Fatal().
				Err(err).
				Msg("Error while creating request signer")
		}
	}

	return c.signer
}

func (c *Container) GetKeyGenerator() signer.KeyGenerator {
	var err error

	publicKey := c.Config.GetString("crypto.ed25519.public")
	privateKey := c.Config.GetString("crypto.ed25519.private")

	_, err = utils.CreateDirectoryFromFile(publicKey, 0o744)

	if err != nil {
		c.GetDefaultLogger().Fatal().
			Str("public_key_path", publicKey).
			Str("private_key_path", privateKey).
			Err(err).
			Msg("Failed to create path to public key")
	}

	_, err = utils.CreateDirectoryFromFile(privateKey, 0o644)

	if err != nil {
		c.GetDefaultLogger().Fatal().
			Str("public_key_path", publicKey).
			Str("private_key_path", privateKey).
			Err(err).
			Msg("Failed to create path to private key")
	}

	generator, err := signer.NewKeyGenerator(privateKey, publicKey)

	if err != nil && !errors.Is(err, signer.ErrKeysAlreadyExist) {
		c.GetDefaultLogger().Fatal().
			Str("public_key_path", publicKey).
			Str("private_key_path", privateKey).
			Err(err).
			Msg("Failed to create ed25519 key generator")
	}

	return generator
}
