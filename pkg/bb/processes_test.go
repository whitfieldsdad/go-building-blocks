package bb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProcesses(t *testing.T) {
	opts := &ProcessOptions{
		IncludeFileMetadata: false,
	}
	ps, err := ListProcesses(opts)
	assert.Nil(t, err)
	assert.NotNil(t, ps)
	assert.True(t, len(ps) > 0)
}

func TestGetProcess(t *testing.T) {
	pid := os.Getpid()
	p, err := GetProcess(pid, nil)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, pid, p.PID)
}
