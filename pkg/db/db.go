package db

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/SSH-Management/server/pkg/db/config"
	"github.com/SSH-Management/server/pkg/db/connector"
	"github.com/SSH-Management/server/pkg/db/drivers/mysql"
)

var ErrNotFound = errors.New("item is not found")

func GetDbConnection(c config.Config) (*gorm.DB, error) {
	var connector connector.Interface
	switch c.Driver {
	case "mysql":
		connector = mysql.New()
	case "sqlite":

	default:
		return nil, fmt.Errorf("Driver %s is not supported", c.Driver)
	}

	dialect, err := connector.Connect(c)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialect, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{},
	})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetConnMaxIdleTime(c.MaxIdleTime)
	sqlDb.SetConnMaxLifetime(c.ConnMaxLifetime)
	sqlDb.SetMaxIdleConns(c.ConnMaxIdle)
	sqlDb.SetMaxOpenConns(c.ConnMaxOpen)

	return db, nil
}
