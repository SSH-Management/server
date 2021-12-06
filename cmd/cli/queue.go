package cli

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/SSH-Management/server/cmd/command"
	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/tasks"
	"github.com/SSH-Management/server/pkg/tasks/processors"
)

func queueWorkerCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "queue:worker",
		Short:             "Run Redis Queue Worker",
		PersistentPreRunE: command.LoadConfig,
		RunE:              handleQueue,
	}
}

func handleQueue(*cobra.Command, []string) error {
	c := command.GetContainer("queue.logging")
	defer c.Close()

	redisOptions := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", c.Config.GetString("redis.host"), c.Config.GetInt("redis.port")),
		Username: c.Config.GetString("redis.username"),
		Password: c.Config.GetString("redis.password"),
		DB:       c.Config.GetInt("redis.queue.db"),
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
		log.Error().
			Err(err).
			Msg("Failed to start Queue Worker")

		return err
	}

	return nil
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
