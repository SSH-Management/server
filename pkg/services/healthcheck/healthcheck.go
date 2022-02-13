package healthcheck

import (
	"context"
	"fmt"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	health_check "google.golang.org/grpc/health/grpc_health_v1"
	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
)

type (
	ClientHealthCheckService struct {
		db     *gorm.DB
		logger *log.Logger
	}
)

func New(db *gorm.DB, logger *log.Logger) *ClientHealthCheckService {
	return &ClientHealthCheckService{
		db:     db,
		logger: logger,
	}
}

func (c *ClientHealthCheckService) Check(ctx context.Context) {
	servers := make([]models.Server, 0, 30)
	c.db.WithContext(ctx).FindInBatches(&servers, 50, func(tx *gorm.DB, batch int) error {
		go func(ctx context.Context, servers []models.Server, db *gorm.DB) {
			for _, server := range servers {
				conn, err := grpc.Dial(fmt.Sprintf("%s:9999", server.IpAddress), grpc.WithTransportCredentials(insecure.NewCredentials()))
				if err != nil {
					c.logger.Error().
						Err(err).
						Str("server_name", server.Name).
						Str("server_ip", server.IpAddress).
						Msg("Cannot connect to server")
					continue
				}

				client := health_check.NewHealthClient(conn)

				res, err := client.Check(ctx, &health_check.HealthCheckRequest{})
				if err != nil {
					c.logger.Error().
						Err(err).
						Str("server_name", server.Name).
						Str("server_ip", server.IpAddress).
						Msg("Cannot run HealthCheck on GRPC")

					continue
				}

				server.Status = mapHealthCheckResponseToString(res.Status)
				result := db.WithContext(ctx).Save(&server)
				if result.Error != nil {
					c.logger.Error().
						Err(err).
						Str("server_name", server.Name).
						Str("server_ip", server.IpAddress).
						Uint64("server_id", server.ID).
						Msg("Cannot update status in Database")
				}

				if err := conn.Close(); err != nil {
					c.logger.Error().
						Err(err).
						Str("server_name", server.Name).
						Str("server_ip", server.IpAddress).
						Msg("Cannot close connection to GRPC server")
				}
			}
		}(ctx, servers, tx)
		return nil
	})
}

func mapHealthCheckResponseToString(status health_check.HealthCheckResponse_ServingStatus) models.ServerStatus {
	switch status {
	case health_check.HealthCheckResponse_NOT_SERVING:
		return models.ServerStatusNotServing
	case health_check.HealthCheckResponse_SERVICE_UNKNOWN:
		return models.ServerStatusUnknown
	case health_check.HealthCheckResponse_SERVING:
		return models.ServerStatusOk
	default:
		return models.ServerStatusUnknown
	}
}
