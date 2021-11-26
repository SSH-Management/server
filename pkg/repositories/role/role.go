package role

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
		Find(context.Context, uint64) (models.Role, error)
		FindByName(context.Context, string) (models.Role, error)
		Create(context.Context) (models.Role, error)
		Delete(context.Context, uint64) error
	}
)

func (r Repository) Find(ctx context.Context, u uint64) (models.Role, error) {
	// TODO implement me
	panic("implement me")
}

func (r Repository) FindByName(ctx context.Context, name string) (models.Role, error) {
	var role models.Role

	result := r.db.WithContext(ctx).
		Where("name = ?", name).
		Limit(1).
		First(&role)

	if result.Error != nil {
		return role, result.Error
	}

	return role, nil
}

func (r Repository) Create(ctx context.Context) (models.Role, error) {
	// TODO implement me
	panic("implement me")
}

func (r Repository) Delete(ctx context.Context, u uint64) error {
	// TODO implement me
	panic("implement me")
}

func New(db *gorm.DB, logger *log.Logger) Repository {
	return Repository{
		db:     db,
		logger: logger,
	}
}
