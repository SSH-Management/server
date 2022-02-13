package container

import (
	"github.com/SSH-Management/server/pkg/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"

	linuxuser "github.com/SSH-Management/linux-user"
	signer "github.com/SSH-Management/request-signer/v4"

	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/repositories/group"
	"github.com/SSH-Management/server/pkg/repositories/role"
	"github.com/SSH-Management/server/pkg/repositories/server"
	userrepo "github.com/SSH-Management/server/pkg/repositories/user"
	"github.com/SSH-Management/server/pkg/services/auth"
	"github.com/SSH-Management/server/pkg/services/password"
	"github.com/SSH-Management/server/pkg/services/user"
)

type Container struct {
	Config *config.Config
	db     *gorm.DB

	defaultLoggerName string

	loggers      map[string]*log.Logger
	redisClients map[int]*redis.Client

	hasher password.Hasher

	userService     user.Interface
	unixUserService linuxuser.UnixInterface

	userRepository   userrepo.Interface
	groupRepository  group.Interface
	serverRepository server.Interface
	roleRepository   role.Interface

	signer signer.Signer

	loginService *auth.LoginService

	validator  *validator.Validate
	translator ut.Translator

	queue   *asynq.Client
	session *session.Store
}

func New(defaultLoggerName string, config *config.Config) *Container {
	return &Container{
		Config:            config,
		defaultLoggerName: defaultLoggerName,
		loggers:           make(map[string]*log.Logger, 1),
		redisClients:      make(map[int]*redis.Client, 16),
	}
}

func (c *Container) Close() error {
	for _, logger := range c.loggers {
		_ = logger.Close()
	}

	for _, client := range c.redisClients {
		_ = client.Close()
	}

	if c.queue != nil {
		_ = c.queue.Close()
	}

	return nil
}
