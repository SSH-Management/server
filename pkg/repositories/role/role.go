package role

import (
	"context"

	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/models"
)

type (
	Repository struct {
		db *gorm.DB
	}

	Interface interface {
		Find(ctx context.Context) ([]models.Role, error)
		FindById(context.Context, uint64) (models.Role, error)
		FindByName(context.Context, string) (models.Role, error)
		Create(context.Context, string, []string) (models.Role, error)
		Delete(context.Context, uint64) error
	}
)

func (r Repository) Find(ctx context.Context) ([]models.Role, error) {
	roles := make([]models.Role, 0, 20)
	result := r.db.WithContext(ctx).Find(&roles)

	if err := result.Error; result.Error != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, db.ErrNotFound
		}

		return nil, err
	}

	return roles, nil
}

func (r Repository) FindById(ctx context.Context, id uint64) (models.Role, error) {
	var role models.Role

	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Limit(1).
		First(&role)

	if err := result.Error; result.Error != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Role{}, db.ErrNotFound
		}

		return models.Role{}, err
	}

	return role, nil
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

func (r Repository) Create(ctx context.Context, name string, perms []string) (models.Role, error) {
	role := models.NewRole(name, perms)

	result := r.db.
		WithContext(ctx).
		Save(role)

	if result.Error != nil {
		return models.Role{}, result.Error
	}

	return role, nil
}

func (r Repository) Delete(ctx context.Context, id uint64) error {
	panic("implement me")
}

func New(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}
