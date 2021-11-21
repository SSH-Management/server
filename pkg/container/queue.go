package container

import (
	"fmt"

	"github.com/hibiken/asynq"
)

func (c *Container) GetQueueClient() *asynq.Client {
	if c.queue == nil {
		c.queue = asynq.NewClient(asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%d", c.Config.GetString("redis.host"), c.Config.GetInt("redis.port")),
			Username: c.Config.GetString("redis.username"),
			Password: c.Config.GetString("redis.password"),
			DB:       5,
		})
	}

	return c.queue
}
