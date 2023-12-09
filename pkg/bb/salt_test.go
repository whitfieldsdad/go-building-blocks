package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSalt(t *testing.T) {
	salt, err := NewSalt(16)
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, salt, "Salt should not be nil")
	assert.Equal(t, 16, len(salt), "Salt length should match")
}
