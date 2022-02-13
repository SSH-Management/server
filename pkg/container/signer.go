package container

import (
	"encoding/binary"
	"errors"
	"os"

	"github.com/SSH-Management/server/pkg/constants"

	signer "github.com/SSH-Management/request-signer/v4"
	"github.com/SSH-Management/utils/v2"
)

func (c *Container) GetSigner() signer.Signer {
	if c.signer != nil {
		return c.signer
	}

	var err error

	generator := c.GetKeyGenerator()

	if err := generator.Generate(); err != nil && !errors.Is(err, signer.ErrKeysAlreadyExist) {
		c.GetDefaultLogger().
			Fatal().
			Err(err).
			Msg("Failed to generate ed25519 keys")
	}

	keysFs := os.DirFS(c.Config.GetString("crypto.ed25519"))

	c.signer, err = signer.NewSignerWithNameAndOrder(keysFs, constants.PrivateKeyFileName, binary.LittleEndian)

	if err != nil {
		c.GetDefaultLogger().
			Fatal().
			Err(err).
			Msg("Error while creating request signer")
	}

	return c.signer
}

func (c *Container) GetKeyGenerator() signer.KeyGenerator {
	var err error

	keysDir := c.Config.GetString("crypto.ed25519")

	_, err = utils.CreateDirectoryFromFile(keysDir, 0o744)

	if err != nil {
		c.GetDefaultLogger().
			Fatal().
			Str("path", keysDir).
			Err(err).
			Msg("Failed to create path to public key")
	}

	generator, err := signer.NewKeyGenerator(keysDir+"/"+constants.PrivateKeyFileName, keysDir+"/"+constants.PublicKeyFileName)

	if err != nil && !errors.Is(err, signer.ErrKeysAlreadyExist) {
		c.GetDefaultLogger().
			Fatal().
			Err(err).
			Msg("Failed to create ed25519 key generator")
	}

	return generator
}
