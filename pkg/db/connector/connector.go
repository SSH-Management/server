package connector

import (
	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/db/config"
)

type Interface interface {
	Connect(v config.Config) (gorm.Dialector, error)
}
