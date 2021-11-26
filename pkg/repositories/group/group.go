package group

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/db"
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
		FindByName(ctx context.Context, name ...string) ([]models.Group, error)
		Create(context.Context, string) (models.Group, error)
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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Group{}, db.ErrNotFound
		}

		return models.Group{}, result.Error
	}

	return group, nil
}

func (r Repository) FindByName(ctx context.Context, name ...string) ([]models.Group, error) {
	groups := make([]models.Group, 0, 20)

	result := r.db.WithContext(ctx).
		Where("name IN ?", name).
		Find(&groups)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, db.ErrNotFound
		}

		return nil, result.Error
	}

	return groups, nil
}

func (r Repository) Create(ctx context.Context, name string) (models.Group, error) {
	g := models.Group{
		Name: name,
	}

	result := r.db.
		WithContext(ctx).
		Create(&g)

	if result.Error != nil {
		return models.Group{}, result.Error
	}

	return g, nil
}

func (r Repository) Delete(ctx context.Context, id uint64) error {
	result := r.db.Delete(&models.Group{Model: models.Model{ID: id}})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return db.ErrNotFound
		}

		return result.Error
	}

	return nil
}

func New(db *gorm.DB, logger *log.Logger) Repository {
	return Repository{
		db:     db,
		logger: logger,
	}
}
