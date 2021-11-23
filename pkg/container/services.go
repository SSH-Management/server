package container

import "github.com/SSH-Management/linux-user"

func (c *Container) GetUnixUserService() user.UnixInterface {
	if c.unixUserService == nil {
		c.unixUserService = user.NewUnixService(c.systemGroups)
	}

	return c.unixUserService
}
