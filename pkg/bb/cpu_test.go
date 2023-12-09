package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCPUs(t *testing.T) {
	cpus, err := ListCPUs()
	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, cpus, "List of CPUs should not be empty")
}
