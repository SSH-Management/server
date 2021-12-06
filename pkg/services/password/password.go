package password

import "errors"

var (
	ErrPasswordMismatch = errors.New("password is invalid")
	ErrInvalidHash      = errors.New("invalid hash")
)

type (
	Hasher interface {
		Hash(string) (string, error)
		Verify(string, string) error
		NeedsRehashing(string) (bool, error)
	}
)
