package user

import (
	"context"
	"testing"

	"github.com/SSH-Management/server/pkg/__mocks__/repositories/group"
	"github.com/SSH-Management/server/pkg/__mocks__/repositories/role"
	"github.com/SSH-Management/server/pkg/db"
	"github.com/SSH-Management/server/pkg/models"

	"github.com/stretchr/testify/require"

	"github.com/SSH-Management/server/pkg/helpers"
)

func TestRepository_Find(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	gormDb, clear := helpers.SetupDatabase()
	defer clear()

	mockRoleRepo := new(role.MockRoleRepository)
	mockGroupRepo := new(group.MockGroupRepository)

	repo := Repository{
		db:        gormDb,
		roleRepo:  mockRoleRepo,
		groupRepo: mockGroupRepo,
	}
	roles := models.GetDefaultRoles()

	result := gormDb.Save(roles)
	assert.NoError(result.Error)

	result = gormDb.Save(&models.User{
		Name:         "Test",
		Surname:      "User",
		Username:     "test_user",
		Email:        "test@test.com",
		Password:     "testpassword123",
		Shell:        "/bin/bash",
		PublicSSHKey: "ssh-ed25519 AAAA",
		RoleID:       roles[0].ID,
	})

	assert.NoError(result.Error)

	users, err := repo.Find(context.Background())

	assert.NoError(err)
	assert.NotNil(users)
	assert.Len(users, 1)
	assert.Equal(uint64(1), users[0].ID)
	assert.Equal("Test", users[0].Name)
}

func TestRepository_FindByEmail_Found(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	gormDb, clear := helpers.SetupDatabase()
	defer clear()

	mockRoleRepo := new(role.MockRoleRepository)
	mockGroupRepo := new(group.MockGroupRepository)

	repo := Repository{
		db:        gormDb,
		roleRepo:  mockRoleRepo,
		groupRepo: mockGroupRepo,
	}

	roles := models.GetDefaultRoles()

	result := gormDb.Save(roles)
	assert.NoError(result.Error)

	result = gormDb.Save(&models.User{
		Name:         "Test",
		Surname:      "User",
		Username:     "test_user",
		Email:        "test@test.com",
		Password:     "testpassword123",
		Shell:        "/bin/bash",
		PublicSSHKey: "ssh-ed25519 AAAA",
		RoleID:       roles[0].ID,
	})

	assert.NoError(result.Error)

	user, err := repo.FindByEmail(context.Background(), "test@test.com")

	assert.NoError(err)
	assert.NotEmpty(user)
	assert.Equal(uint64(1), user.ID)
	assert.Equal("Test", user.Name)
}

func TestRepository_FindByEmail_NotFound(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	gormDb, clear := helpers.SetupDatabase()
	defer clear()

	mockRoleRepo := new(role.MockRoleRepository)
	mockGroupRepo := new(group.MockGroupRepository)

	repo := Repository{
		db:        gormDb,
		roleRepo:  mockRoleRepo,
		groupRepo: mockGroupRepo,
	}

	user, err := repo.FindByEmail(context.Background(), "test@test.com")

	assert.Error(err)
	assert.Empty(user)
	assert.ErrorIs(err, db.ErrNotFound)
}
