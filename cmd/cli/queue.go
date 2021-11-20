package cli

import (
	"github.com/SSH-Management/server/pkg/container"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


func queueWorkerCommand(c *container.Container, v *viper.Viper) *cobra.Command {
	return &cobra.Command{
		Use: "queue:worker",
		Short: "Run Redis Queue Worker",
		Run: handleQueue(c, v),
	}
}


func handleQueue(c *container.Container, v *viper.Viper) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {}
}
