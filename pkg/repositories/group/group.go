package group

import (
	"context"

	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
)

type (
	Repository struct {
		db     *gorm.DB
		logger *log.Logger
	}

	Interface interface {
		Find(context.Context, uint64) (models.Group, error)
		FindByName(context.Context, string) (models.Group, error)
		Create(context.Context) (models.Group, error)
		Delete(context.Context, uint64) error
	}
)

func (r Repository) Find(ctx context.Context, id uint64) (models.Group, error) {
	var group models.Group

	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Limit(1).
		First(&group)

	if result.Error != nil {
		return group, result.Error
	}

	return group, nil
}

func (r Repository) FindByName(ctx context.Context, name string) (models.Group, error) {
	var group models.Group

	result := r.db.WithContext(ctx).
		Where("name = ?", name).
		Limit(1).
		First(&group)

	if result.Error != nil {
		return group, result.Error
	}

	return group, nil
}

func (r Repository) Create(ctx context.Context) (models.Group, error) {
	panic("implement me")
}

func (r Repository) Delete(ctx context.Context, u uint64) error {
	panic("implement me")
}

func New(db *gorm.DB, logger *log.Logger) Repository {
	return Repository{
		db:     db,
		logger: logger,
	}
}
