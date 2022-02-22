package password

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/SSH-Management/utils/v2"
)

var (
	_ Hasher = BcryptPassword{}
)

type BcryptPassword struct {
	cost int
}

func NewBcrypt(cost int) BcryptPassword {
	if cost < 7 || cost > 32 {
		panic(fmt.Sprintf("Cost: %d must be betweeen 8 and 32", cost))
	}

	return BcryptPassword{
		cost: cost,
	}
}

func (b BcryptPassword) Hash(password string) (string, error) {
	passwordBytes := utils.UnsafeBytes(password)

	hash, err := bcrypt.GenerateFromPassword(passwordBytes, b.cost)
	if err != nil {
		return "", err
	}

	return utils.UnsafeString(hash), err
}

func (b BcryptPassword) Verify(hash, password string) error {
	hashBytes := utils.UnsafeBytes(hash)
	passwordBytes := utils.UnsafeBytes(password)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)

	if err == nil {
		return nil
	}

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return ErrPasswordMismatch
	}

	return ErrInvalidHash
}

func (b BcryptPassword) NeedsRehashing(hash string) (bool, error) {
	cost, err := bcrypt.Cost(utils.UnsafeBytes(hash))
	if err != nil {
		return false, ErrInvalidHash
	}

	if b.cost != cost {
		return true, nil
	}

	return false, nil
}
