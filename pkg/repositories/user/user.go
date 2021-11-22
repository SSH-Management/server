package user

import (
	"context"
	"time"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/repositories/role"

	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
)

type (
	Repository struct {
		db     *gorm.DB
		logger *log.Logger

		roleRepo role.Interface
	}

	Interface interface {
		FindByGroup(ctx context.Context, id uint64) ([]models.User, error)
		Create(ctx context.Context, dto dto.User) (models.User, error)
		Delete(context.Context, uint64) error
	}
)

func New(db *gorm.DB, logger *log.Logger, roleRepo role.Interface) Repository {
	return Repository{
		db:       db,
		logger:   logger,
		roleRepo: roleRepo,
	}
}

func (r Repository) FindByGroup(ctx context.Context, id uint64) ([]models.User, error) {
	users := make([]models.User, 0, 10)

	result := r.db.
		WithContext(ctx).
		Preload("Groups", "id = ?", id).
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r Repository) Create(ctx context.Context, dto dto.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	roleModel, err := r.roleRepo.FindByName(ctx, dto.Role)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:     dto.Name,
		Surname:  dto.Surname,
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		Shell:    dto.Shell,
		RoleID:   roleModel.ID,
	}

	result := r.db.
		WithContext(ctx).
		Create(&user)

	err = result.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, db.ErrNotFound
		}
		return models.User{}, err
	}

	return user, nil
}

func (r Repository) Delete(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	result := r.db.
		WithContext(ctx).
		Model(&models.User{}).
		Delete(id)

	return result.Error
}
