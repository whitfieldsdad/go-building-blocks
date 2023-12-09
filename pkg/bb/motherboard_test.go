package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMotherboard(t *testing.T) {
	mobo, err := GetMotherboard()
	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, mobo, "Motherboard information should be available")
}
