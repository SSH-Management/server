package container

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/db/config"
)

func (c *Container) GetDbConnection() *gorm.DB {
	if c.db == nil {
		var err error
		var cfg config.Config

		err = c.Config.Sub("database").Unmarshal(&cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while reading DB Config")
		}

		c.db, err = db.GetDbConnection(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while connecting to MySQL database")
		}
	}

	return c.db
}
