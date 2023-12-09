package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListMemoryModules(t *testing.T) {
	rams, err := ListMemoryModules()
	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, rams, "List of memory modules should not be empty")
}
