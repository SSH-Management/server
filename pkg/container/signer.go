package container

import (
	"github.com/SSH-Management/request-signer/v2"
	"github.com/SSH-Management/utils"
)

func (c *Container) GetSigner() signer.Interface {
	if c.signer == nil {
		var err error

		publicKey := c.Config.GetString("crypto.ed25519.public")
		privateKey := c.Config.GetString("crypto.ed25519.private")

		_, err = utils.CreatePath(publicKey, 0644)

		if err != nil {
			c.Logger.Fatal().
				Str("public_key_path", publicKey).
				Str("private_key_path", privateKey).
				Err(err).
				Msg("Failed to create path to keys")
		}

		c.signer, err = signer.New(publicKey, privateKey)

		if err != nil {
			c.Logger.Fatal().
				Str("public_key_path", publicKey).
				Str("private_key_path", privateKey).
				Err(err).
				Msg("Error while creating request signer")
		}
	}

	return c.signer
}
