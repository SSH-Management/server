package server

import (
	"context"
	"database/sql"
	"errors"

	"gorm.io/gorm"

	"github.com/SSH-Management/protobuf/server/clients"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
	"github.com/SSH-Management/server/pkg/repositories/group"
)

var _ Interface = &Repository{}

type (
	Repository struct {
		db     *gorm.DB
		logger *log.Logger

		groupRepo group.Interface
	}

	Interface interface {
		FindAll(context.Context) ([]models.Server, error)
		FindByPrivateIP(context.Context, string) (models.Server, error)
		Find(context.Context, uint64) (models.Server, error)
		Create(context.Context, *clients.CreateClientRequest) (models.Server, error)
		Delete(context.Context, uint64) error
	}
)

func (r Repository) FindAll(ctx context.Context) ([]models.Server, error) {
	servers := make([]models.Server, 0, 100)

	result := r.db.Find(&servers)

	if result.Error != nil {
		return nil, result.Error
	}

	return servers, nil
}

func (r Repository) Find(ctx context.Context, id uint64) (models.Server, error) {
	panic("implement me")
}

func (r Repository) FindByPrivateIP(ctx context.Context, ip string) (models.Server, error) {
	var s models.Server

	result := r.db.
		WithContext(ctx).
		Where("ip = ?", ip).
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

func (r Repository) createGroupIfNotExists(ctx context.Context, name string) (models.Group, error) {
	var g models.Group
	groups, err := r.groupRepo.FindByName(ctx, name)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			g, err = r.groupRepo.Create(ctx, name)

			if err != nil {
				return g, err
			}
		} else {
			return g, err
		}
	}

	if len(groups) == 0 {
		g, err = r.groupRepo.Create(ctx, name)

		if err != nil {
			return g, err
		}

		return g, nil
	}

	return groups[0], nil
}

func (r Repository) Create(ctx context.Context, dto *clients.CreateClientRequest) (models.Server, error) {
	publicIpSql := sql.NullString{
		String: "",
		Valid:  false,
	}

	g, err := r.createGroupIfNotExists(ctx, dto.Group)
	if err != nil {
		return models.Server{}, err
	}

	server := models.Server{
		Name:            dto.Name,
		IpAddress:       dto.Ip,
		PublicIpAddress: publicIpSql,
		GroupID:         g.ID,
		Status:          models.ServerStatusUnknown,
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
	result := r.db.
		WithContext(ctx).
		Delete(&models.Server{Model: models.Model{ID: id}})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return db.ErrNotFound
		}

		return result.Error
	}

	return nil
}

func New(db *gorm.DB, logger *log.Logger, groupRepo group.Interface) Repository {
	return Repository{
		db:        db,
		logger:    logger,
		groupRepo: groupRepo,
	}
}
