package container

import (
	"github.com/SSH-Management/server/pkg/repositories/group"
	"github.com/SSH-Management/server/pkg/repositories/role"
	"github.com/SSH-Management/server/pkg/repositories/server"
)

func (c *Container) GetGroupRepository() group.Interface {
	if c.groupRepository == nil {
		c.groupRepository = group.New(c.GetDbConnection())
	}

	return c.groupRepository
}

func (c *Container) GetServerRepository() server.Interface {
	if c.serverRepository == nil {
		c.serverRepository = server.New(
			c.GetDbConnection(),
			c.GetDefaultLogger(),
			c.GetGroupRepository(),
		)
	}

	return c.serverRepository
}

func (c *Container) GetRoleRepository() role.Interface {
	if c.roleRepository == nil {
		c.roleRepository = role.New(c.GetDbConnection())
	}

	return c.roleRepository
}
