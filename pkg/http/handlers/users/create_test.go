package users_test

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"testing"
	"time"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	linuxuser "github.com/SSH-Management/linux-user"
	"github.com/SSH-Management/server/pkg/__mocks__/services/user"
	"github.com/SSH-Management/server/pkg/dto"
	"github.com/SSH-Management/server/pkg/helpers"
	"github.com/SSH-Management/server/pkg/http/handlers/users"
	"github.com/SSH-Management/server/pkg/log"
	"github.com/SSH-Management/server/pkg/models"
	"github.com/SSH-Management/server/pkg/repositories/group"
	"github.com/SSH-Management/server/pkg/repositories/role"
	userrepo "github.com/SSH-Management/server/pkg/repositories/user"
	"github.com/SSH-Management/server/pkg/services/password"
	realuserservice "github.com/SSH-Management/server/pkg/services/user"
)

var userSuccessDto = dto.CreateUser{
	User: dto.User{
		Name:         "Test",
		Surname:      "Test",
		Email:        "test@test.com",
		Username:     "test_user",
		Password:     "password123",
		Shell:        "/bin/bash",
		Role:         "Admin",
		Groups:       []string{"test_group"},
		SystemGroups: []string{"sudo"},
	},
	PublicSSHKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIK0wmN/Cr3JXqmLW7u+g9pTh+wyqDHpSQEIQczXkVx9q test@test.com",
}

func TestCreateUserHandler_Unit(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	mockUserService := new(user.MockUserService)

	app, _ := helpers.CreateApplication()
	validator, _ := helpers.GetValidator()

	app.Post("/", users.CreateUserHandler(mockUserService, validator))

	userModel := models.User{
		Model: models.Model{
			ID:        1,
			CreatedAt: time.Now(),
		},
		Name:         "Test",
		Surname:      "Test",
		Username:     "test_user",
		Email:        "test@test.com",
		Password:     "hashed_password",
		Shell:        "/bin/bash",
		PublicSSHKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIK0wmN/Cr3JXqmLW7u+g9pTh+wyqDHpSQEIQczXkVx9q test@test.com",
		Role:         models.Role{Model: models.Model{ID: 1}, Name: "Admin"},
		RoleID:       1,
		Groups: []models.Group{
			{Model: models.Model{ID: 1}, Name: "test_groups", IsSystemGroup: false},
			{Model: models.Model{ID: 2}, Name: "sudo", IsSystemGroup: true},
		},
	}

	mockUserService.On("Create", mock.Anything, userSuccessDto).
		Once().
		Return(userModel, []byte{1}, nil)

	res := helpers.Post(app, "/", helpers.WithBody(userSuccessDto))

	assert.Equal(http.StatusCreated, res.StatusCode)

	mockUserService.AssertExpectations(t)
}

func TestCreateUserHandler_Success(t *testing.T) {
	t.Parallel()

	if runtime.GOOS != "linux" {
		t.Skipf("Cannot run this test under %s", runtime.GOOS)
	}

	assert := require.New(t)

	db, clean := helpers.SetupDatabase()

	defer clean()
	app, _ := helpers.CreateApplication()
	validator, _ := helpers.GetValidator()

	roleModel := models.NewRole(userSuccessDto.User.Role, []string{})
	db.Create(&roleModel)

	db.Create(&models.Group{
		Name:          "test_group",
		IsSystemGroup: false,
	})

	db.Create(&models.Group{
		Name:          "sudo",
		IsSystemGroup: true,
	})

	userRepo := userrepo.New(db, role.New(db), group.New(db))

	logger, _ := log.NewTest(t, zerolog.ErrorLevel)

	var unixService linuxuser.UnixInterface

	unixService = linuxuser.NewUnixService(map[string]string{"sudo": "sudo"})

	redisHost, redisPort, redisUsername, redisPassword := helpers.GetRedisConnParams()
	client := helpers.GetRedisConnection(0)
	defer client.FlushDBAsync(context.Background())
	userService := realuserservice.New(
		userRepo,
		unixService,
		logger,
		asynq.NewClient(asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
			Username: redisUsername,
			Password: redisPassword,
			DB:       0,
		}),
		password.NewBcrypt(8),
	)

	app.Post("/", users.CreateUserHandler(userService, validator))

	res := helpers.Post(app, "/", helpers.WithBody(userSuccessDto))

	assert.Equal(http.StatusCreated, res.StatusCode)
}
