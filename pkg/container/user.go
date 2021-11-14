package container

import (
	userrepo "github.com/SSH-Management/server/pkg/repositories/user"
	"github.com/SSH-Management/server/pkg/services/user"
)

func (c *Container) GetUserService() user.Interface {
	if c.userService == nil {
		c.userService = user.New(c.GetUserRepository(), c.GetUnixUserService(), c.Logger)
	}

	return c.userService
}

func (c *Container) GetUserRepository() userrepo.Interface {
	if c.userRepository == nil {
		c.userRepository = userrepo.New(
			c.GetDbConnection(),
			c.Logger,
			c.GetRoleRepository(),
		)
	}

	return c.userRepository
}
