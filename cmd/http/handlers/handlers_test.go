package handlers_test

import (
	"testing"

	"github.com/SSH-Management/server/cmd/http/handlers"

	"github.com/stretchr/testify/require"
)

func TestRedirect(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	url, err := handlers.Redirect("/")

	assert.NoError(err)
	assert.Equal("/?message=Message+sent&status=success", url)
}
