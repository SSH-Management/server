package server

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	sdk "github.com/SSH-Management/server-sdk"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
	"github.com/SSH-Management/server/pkg/repositories/group"
)

type (
	Repository struct {
		db     *gorm.DB
		logger *log.Logger

		groupRepo group.Interface
	}

	Interface interface {
		FindByPrivateIP(context.Context, string) (models.Server, error)
		Find(context.Context, uint64) (models.Server, error)
		Create(context.Context, sdk.NewClientRequest) (models.Server, error)
		Delete(context.Context, uint64) error
	}
)

func (r Repository) Find(ctx context.Context, id uint64) (models.Server, error) {
	panic("implement me")
}

func (r Repository) FindByPrivateIP(ctx context.Context, ip string) (models.Server, error) {
	var s models.Server

	result := r.db.
		WithContext(ctx).
		Where("private_ip = ?", ip).
		Limit(1).
		First(&s)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return models.Server{}, db.ErrNotFound
		}

		return models.Server{}, result.Error
	}

	return s, nil
}

func (r Repository) Create(ctx context.Context, dto sdk.NewClientRequest) (models.Server, error) {
	publicIpSql := sql.NullString{
		String: dto.PublicIp,
		Valid:  dto.PublicIp != "",
	}

	g, err := r.groupRepo.FindByName(ctx, dto.Group)

	if err != nil {
		return models.Server{}, err
	}

	server := models.Server{
		Name:            dto.Name,
		IpAddress:       dto.Ip,
		PublicIpAddress: publicIpSql,
		GroupID:         g.ID,
	}

	result := r.db.
		WithContext(ctx).
		Create(&server)

	if result.Error != nil {
		return models.Server{}, result.Error
	}

	return server, nil
}

func (r Repository) Delete(ctx context.Context, id uint64) error {
	panic("implement me")
}

func New(db *gorm.DB, logger *log.Logger, groupRepo group.Interface) Repository {
	return Repository{
		db:        db,
		logger:    logger,
		groupRepo: groupRepo,
	}
}
