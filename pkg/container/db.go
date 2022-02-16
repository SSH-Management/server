package container

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
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

func (c *Container) GetRedisClient(db int) *redis.Client {
	if client, ok := c.redisClients[db]; ok {
		return client
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Config.GetString("redis.host"), c.Config.GetInt("redis.port")),
		Password: c.Config.GetString("redis.password"),
		Username: c.Config.GetString("redis.username"),
		DB:       db,
	})

	c.redisClients[db] = rdb

	return rdb
}

func (c *Container) GetRawSQL() *sql.DB {
	sqlDB, err := c.GetDbConnection().DB()
	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Cannot obtain instance of *sql.DB")
	}

	return sqlDB
}
