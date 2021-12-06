package container

import "github.com/SSH-Management/server/pkg/services/auth"

func (c *Container) GetLoginService() *auth.LoginService {
	if c.loginService == nil {
		c.loginService = auth.NewLoginService(
			c.GetUserRepository(),
			c.GetPasswordHasher(),
		)
	}

	return c.loginService
}
