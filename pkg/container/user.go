package container

import (
	userrepo "github.com/SSH-Management/server/pkg/repositories/user"
	"github.com/SSH-Management/server/pkg/services/user"
)

func (c *Container) GetUserService() user.Interface {
	if c.userService == nil {
		s := user.New(
			c.GetUserRepository(),
			c.GetUnixUserService(),
			c.GetDefaultLogger(),
			c.GetQueueClient(),
			c.GetPasswordHasher(),
		)
		c.userService = s
	}

	return c.userService
}

func (c *Container) GetUserRepository() userrepo.Interface {
	if c.userRepository == nil {
		c.userRepository = userrepo.New(
			c.GetDbConnection(),
			c.GetDefaultLogger(),
			c.GetRoleRepository(),
			c.GetGroupRepository(),
		)
	}

	return c.userRepository
}
