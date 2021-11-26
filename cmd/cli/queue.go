package cli

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/tasks"
	"github.com/SSH-Management/server/pkg/tasks/processors"
)

func queueWorkerCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "queue:worker",
		Short:             "Run Redis Queue Worker",
		PersistentPreRunE: loadConfig,
		Run:               handleQueue(),
	}
}

func handleQueue() func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		c := getContainer("queue.logging")
		defer c.Close()

		redisOptions := asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%d", c.Config.GetString("redis.host"), c.Config.GetInt("redis.port")),
			Username: c.Config.GetString("redis.username"),
			Password: c.Config.GetString("redis.password"),
			DB:       5,
		}

		srv := asynq.NewServer(
			redisOptions,
			asynq.Config{
				Concurrency: c.Config.GetInt("queue.concurrency"),
				Queues: map[string]int{
					"crititcal": 6,
					"default":   3,
				},
			},
		)

		mux := asynq.NewServeMux()

		registerQueueHandlers(c, mux)

		if err := srv.Run(mux); err != nil {
			log.Fatal().Err(err).Msg("Failed to start Queue Worker")
		}
	}
}

func registerQueueHandlers(c *container.Container, mux *asynq.ServeMux) {
	mux.Handle(tasks.TypeNewUser, processors.NewUserCreatedProcessor(
		c.GetDbConnection(),
		c.GetQueueClient(),
		c.GetDefaultLogger(),
	))

	mux.Handle(tasks.TypeNotifyServerForNewUser, processors.NewNotifyServerForNewUser(
		c.GetDefaultLogger(),
	))
}
