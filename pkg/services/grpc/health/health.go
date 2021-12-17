package health

import (
	"context"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	health_check "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var _ health_check.HealthServer = &HealthCheckService{}

type HealthCheckService struct {
	db    *gorm.DB
	redis *redis.Client
}

func New(db *gorm.DB, redisClient *redis.Client) *HealthCheckService {
	return &HealthCheckService{
		db:    db,
		redis: redisClient,
	}
}

func (h *HealthCheckService) Check(ctx context.Context, req *health_check.HealthCheckRequest) (*health_check.HealthCheckResponse, error) {
	result := h.redis.Ping(ctx)

	if err := result.Err(); err != nil {
		return &health_check.HealthCheckResponse{
			Status: health_check.HealthCheckResponse_NOT_SERVING,
		}, err
	}

	db, err := h.db.DB()
	if err != nil {
		return &health_check.HealthCheckResponse{
			Status: health_check.HealthCheckResponse_NOT_SERVING,
		}, err
	}

	if err := db.PingContext(ctx); err != nil {
		return &health_check.HealthCheckResponse{
			Status: health_check.HealthCheckResponse_NOT_SERVING,
		}, err
	}

	return &health_check.HealthCheckResponse{
		Status: health_check.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthCheckService) Watch(req *health_check.HealthCheckRequest, server health_check.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Not implemented")
}
