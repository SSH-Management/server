package user

import (
	"context"
	"errors"
	"time"

	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/repositories/group"
	"github.com/SSH-Management/server/pkg/repositories/role"

	"gorm.io/gorm"

	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
)

var _ Interface = &Repository{}

type (
	Repository struct {
		db     *gorm.DB
		logger *log.Logger

		roleRepo  role.Interface
		groupRepo group.Interface
	}

	Interface interface {
		Find(context.Context) ([]models.User, error)
		FindByGroup(context.Context, uint64) ([]models.User, error)
		FindByEmail(context.Context, string) (models.User, error)
		Create(context.Context, dto.User, string) (models.User, error)
		Delete(context.Context, uint64) error
	}
)

func New(db *gorm.DB, logger *log.Logger, roleRepo role.Interface, groupRepo group.Interface) Repository {
	return Repository{
		db:        db,
		logger:    logger,
		roleRepo:  roleRepo,
		groupRepo: groupRepo,
	}
}

func (r Repository) Find(ctx context.Context) ([]models.User, error) {
	users := make([]models.User, 0, 10)

	result := r.db.Find(&users)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, db.ErrNotFound
		}

		return nil, result.Error
	}

	return users, nil
}

func (r Repository) FindByGroup(ctx context.Context, id uint64) ([]models.User, error) {
	users := make([]models.User, 0, 10)

	result := r.db.WithContext(ctx).
		Model(&models.User{}).
		Joins("inner join user_groups on users.id = user_groups.user_id").
		Where("user_groups.group_id = ?", id).
		Select("users.*").
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r Repository) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var u models.User

	result := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		Limit(1).
		First(&u)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, db.ErrNotFound
		}

		return models.User{}, result.Error
	}

	return u, nil
}

func (r Repository) Create(ctx context.Context, dto dto.User, publicKey string) (models.User, error) {
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()

	roleModel, err := r.roleRepo.FindByName(ctx, dto.Role)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:         dto.Name,
		Surname:      dto.Surname,
		Username:     dto.Username,
		Email:        dto.Email,
		Password:     dto.Password,
		Shell:        dto.Shell,
		RoleID:       roleModel.ID,
		PublicSSHKey: publicKey,
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

	groups, err := r.groupRepo.FindByName(ctx, dto.Groups...)
	if err != nil {
		if err := r.Delete(ctx, user.ID); err != nil {
			return models.User{}, err
		}

		return models.User{}, err
	}

	err = r.db.
		WithContext(ctx).
		Model(&user).
		Association("Groups").
		Append(groups)

	if err != nil {
		if err := r.Delete(ctx, user.ID); err != nil {
			return models.User{}, err
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
