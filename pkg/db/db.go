package db

import (
	"errors"
	"fmt"

	"github.com/SSH-Management/server/pkg/db/config"
	"github.com/SSH-Management/server/pkg/db/connector"
	"github.com/SSH-Management/server/pkg/db/drivers/mysql"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("item is not found")

func GetDbConnection(c config.Config) (*gorm.DB, error) {
	var dbConnector connector.Interface
	switch c.Driver {
	case "mysql":
		dbConnector = mysql.New()
	default:
		return nil, fmt.Errorf("driver '%s' is not supported", c.Driver)
	}

	dialect, err := dbConnector.Connect(c)

	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialect, &gorm.Config{})

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
