package bb

import (
	"crypto/rand"

	"github.com/pkg/errors"
)

func NewSalt(len int) ([]byte, error) {
	salt := make([]byte, len)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate salt")
	}
	return salt, nil
}
