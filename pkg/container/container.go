package container

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	linux_user "github.com/SSH-Management/linux-user"
	"github.com/SSH-Management/request-signer/v3"

	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/repositories/group"
	"github.com/SSH-Management/server/pkg/repositories/role"
	"github.com/SSH-Management/server/pkg/repositories/server"
	userrepo "github.com/SSH-Management/server/pkg/repositories/user"
	"github.com/SSH-Management/server/pkg/services/user"
)

type Container struct {
	systemGroups map[string]string
	Logger       *log.Logger
	Config       *viper.Viper
	db           *gorm.DB

	userService     user.Interface
	unixUserService linux_user.UnixInterface

	userRepository   userrepo.Interface
	groupRepository  group.Interface
	serverRepository server.Interface
	roleRepository   role.Interface

	signer signer.Signer

	validator  *validator.Validate
	translator ut.Translator
}

func New(config *viper.Viper) (*Container, error) {
	logger, err := log.New(
		config.GetString("logging.file"),
		config.GetString("logging.level"),
		config.GetBool("logging.console"),
		config.GetUint32("logging.sample"),
	)

	if err != nil {
		return nil, err
	}

	return &Container{
		Logger:       logger,
		Config:       config,
		systemGroups: config.GetStringMapString("system_groups"),
	}, nil
}

func (c *Container) Close() error {
	if c.Logger != nil {
		_ = c.Logger.Close()
	}

	return nil
}
