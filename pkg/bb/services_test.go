package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListServices(t *testing.T) {
	opts := &ProcessOptions{
		IncludeFileMetadata: false,
	}
	_, err := ListServices(opts)
	assert.Nil(t, err)
}
