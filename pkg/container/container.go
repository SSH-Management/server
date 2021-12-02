package container

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/hibiken/asynq"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	linux_user "github.com/SSH-Management/linux-user"
	signer "github.com/SSH-Management/request-signer/v3"

	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/repositories/group"
	"github.com/SSH-Management/server/pkg/repositories/role"
	"github.com/SSH-Management/server/pkg/repositories/server"
	userrepo "github.com/SSH-Management/server/pkg/repositories/user"
	password "github.com/SSH-Management/server/pkg/services/password"
	"github.com/SSH-Management/server/pkg/services/user"
)

type Container struct {
	systemGroups map[string]string
	Config       *viper.Viper
	db           *gorm.DB

	defaultLoggerName string

	loggers map[string]*log.Logger

	hasher password.Hasher

	userService     user.Interface
	unixUserService linux_user.UnixInterface

	userRepository   userrepo.Interface
	groupRepository  group.Interface
	serverRepository server.Interface
	roleRepository   role.Interface

	signer signer.Signer

	validator  *validator.Validate
	translator ut.Translator

	queue *asynq.Client
}

func New(defaultLoggerName string, config *viper.Viper) *Container {
	return &Container{
		Config:            config,
		defaultLoggerName: defaultLoggerName,
		loggers:           make(map[string]*log.Logger, 1),
		systemGroups:      config.GetStringMapString("system_groups"),
	}
}

func (c *Container) Close() error {
	for _, logger := range c.loggers {
		_ = logger.Close()
	}

	if c.queue != nil {
		_ = c.queue.Close()
	}

	return nil
}
