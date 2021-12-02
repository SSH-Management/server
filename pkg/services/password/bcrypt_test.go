package password

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

const (
	Prefix = "$2a$10$"
)

func TestBcryptHashing(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	hash, err := bcrypt.GenerateFromPassword([]byte("test_password"), bcrypt.DefaultCost)

	assert.NoError(err)
	assert.True(strings.HasPrefix(string(hash), Prefix))

	b := BcryptPassword{cost: bcrypt.DefaultCost}

	otherHash, err := b.Hash("test_password")

	assert.NoError(err)
	assert.NotEqual(hash, otherHash)
	assert.True(strings.HasPrefix(otherHash, Prefix))
}

func TestBcryptVerify(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	hash, err := bcrypt.GenerateFromPassword([]byte("test_password"), bcrypt.DefaultCost)

	assert.NoError(err)

	t.Run("ValidPassword", func(t *testing.T) {
		b := BcryptPassword{cost: bcrypt.DefaultCost}

		err := b.Verify(string(hash), "test_password")

		assert.NoError(err)
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		b := BcryptPassword{cost: bcrypt.DefaultCost}

		err := b.Verify(string(hash), "invalid-password-test")

		assert.Error(err)
		assert.ErrorIs(ErrPasswordMismatch, err)
	})

	t.Run("InValidHash", func(t *testing.T) {
		b := BcryptPassword{cost: bcrypt.DefaultCost}

		err := b.Verify("hjashdjkjashdlkjasd", "invalid-password-test")

		assert.Error(err)
		assert.ErrorIs(ErrInvalidHash, err)
	})
}
