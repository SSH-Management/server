package container

import (
	"strings"

	user "github.com/SSH-Management/linux-user"
	"github.com/SSH-Management/server/pkg/log"

	password "github.com/SSH-Management/server/pkg/services/password"
)

func (c *Container) GetUnixUserService() user.UnixInterface {
	if c.unixUserService == nil {
		c.unixUserService = user.NewUnixService(
			c.systemGroups,
			log.UnixServiceLogger{
				Logger: c.GetDefaultLogger(),
			},
		)
	}

	return c.unixUserService
}


func (c *Container) GetPasswordHasher() password.Hasher {
	if c.hasher == nil {
		driver := c.Config.GetString("crypto.password.driver")

		switch strings.ToLower(driver) {
		case "bcrypt":
			c.hasher = password.NewBcrypt(c.Config.GetInt("crypto.password.bcrypt.cost"))
		default:
			c.GetDefaultLogger().
				Fatal().
				Str("driver", driver).
				Msg("Invalid password hashing algorithm")
		}
	}

	return c.hasher
}
