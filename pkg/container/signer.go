package container

import (
	signer "github.com/SSH-Management/request-signer"
)

func (c *Container) GetSigner() signer.Interface {
	if c.signer == nil {
		var err error

		publicKey := c.Config.GetString("crypto.ed25519.public")
		privateKey := c.Config.GetString("crypto.ed25519.private")

		c.signer, err = signer.NewSigner(publicKey, privateKey)

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
