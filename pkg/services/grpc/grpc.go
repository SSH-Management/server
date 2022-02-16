package grpc

import (
	pb_client "github.com/SSH-Management/protobuf/server/clients"
	"google.golang.org/grpc"
	health_check "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/SSH-Management/server/pkg/container"
	"github.com/SSH-Management/server/pkg/services/grpc/client"
	"github.com/SSH-Management/server/pkg/services/grpc/health"
)

func RegisterServices(c *container.Container, server grpc.ServiceRegistrar) {
	serverPublicKey, err := c.Config.LoadServerPublicSSHKey()
	if err != nil {
		c.GetDefaultLogger().
			Fatal().
			Err(err).
			Msg("Failed to load Server's Public Key")
	}

	health_check.RegisterHealthServer(
		server,
		health.New(
			c.GetDbConnection(),
			c.GetRedisClient(0),
		),
	)

	pb_client.RegisterClientServiceServer(
		server,
		client.New(
			serverPublicKey,
			c.GetDefaultLogger(),
			c.GetServerRepository(),
			c.GetUserRepository(),
			c.GetValidator(),
		),
	)
}
