package groups

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	mockgroup "github.com/SSH-Management/server/pkg/__mocks__/repositories/group"
	"github.com/SSH-Management/server/pkg/helpers"
	"github.com/SSH-Management/server/pkg/models"
	"github.com/SSH-Management/server/pkg/repositories/group"
)

func TestGetGroups_Error(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _ := helpers.CreateApplication()

	mockGroupRepo := new(mockgroup.MockGroupRepository)

	app.Get("/", GetGroups(mockGroupRepo))

	mockGroupRepo.On("FindById", mock.Anything, false).
		Once().
		Return(nil, errors.New("some db error"))

	res := helpers.Get(app, "/?type=server")

	assert.Equal(http.StatusInternalServerError, res.StatusCode)
	mockGroupRepo.AssertExpectations(t)
}

func TestGetGroups_InvalidType(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _ := helpers.CreateApplication()

	app.Get("/", GetGroups(group.New(nil)))

	res := helpers.Get(app, "/?type=invalid_type")

	assert.Equal(http.StatusUnprocessableEntity, res.StatusCode)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	bytes, _ := ioutil.ReadAll(res.Body)
	assert.Equal([]byte("Type not valid for Groups: system or user"), bytes)
}

func TestGetGroups_Integration(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	app, _ := helpers.CreateApplication()
	db, clean := helpers.SetupDatabase()

	defer clean()

	db.Save([]models.Group{
		{Name: "test_server_g1", IsSystemGroup: false},
		{Name: "test_system_g1", IsSystemGroup: true},
	})

	app.Get("/", GetGroups(group.New(db)))

	t.Run("SystemGroups", func(t *testing.T) {
		res := helpers.Get(app, "/?type=system")

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Equal(fiber.MIMEApplicationJSON, res.Header.Get("Content-Type"))

		var groups []models.Group
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)
		assert.NoError(json.NewDecoder(res.Body).Decode(&groups))
		assert.Len(groups, 1)
		assert.Equal("test_system_g1", groups[0].Name)
	})

	t.Run("UserGroups", func(t *testing.T) {
		res := helpers.Get(app, "/?type=server")

		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Equal(fiber.MIMEApplicationJSON, res.Header.Get("Content-Type"))

		var groups []models.Group
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)
		assert.NoError(json.NewDecoder(res.Body).Decode(&groups))
		assert.Len(groups, 1)
		assert.Equal("test_server_g1", groups[0].Name)
	})
}
