package db

import (
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/db/config"
)

var ErrNotFound = errors.New("item is not found")


func createPostgresDatabaseConnection(c config.Config) gorm.Dialector {
	return postgres.New(postgres.Config{
		DSN:                  c.FormatConnectionString(),
		PreferSimpleProtocol: false,
		WithoutReturning:     false,
	})
}

func GetDbConnection(c config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(createPostgresDatabaseConnection(c), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
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
