package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentUser(t *testing.T) {
	user, err := GetCurrentUser()
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Username)
}

func TestGetUser(t *testing.T) {
	a, err := GetCurrentUser()
	assert.Nil(t, err)
	assert.NotNil(t, a)

	b, err := GetUser(a.Username)
	assert.Nil(t, err)
	assert.NotNil(t, b)

	assert.Equal(t, a, b)

	// Ensure that all required fields are populated.
	assert.NotEmpty(t, b.Name)
	assert.NotEmpty(t, b.Username)
	assert.NotEmpty(t, b.UID)
	assert.NotEmpty(t, b.GID)
	assert.NotEmpty(t, b.GroupIds)
	assert.NotEmpty(t, b.HomeDir)
}
