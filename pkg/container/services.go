package container

import (
	user "github.com/SSH-Management/linux-user"
	"github.com/SSH-Management/server/pkg/log"
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
