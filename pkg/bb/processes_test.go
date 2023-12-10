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

func TestGetProcessUUID(t *testing.T) {
	u, err := GetProcessUUID(123, 456)
	assert.Nil(t, err)
	assert.Equal(t, "17e3b0a9-32bf-5f5f-b0c9-2b9846033763", u)
}

func TestGetProcess(t *testing.T) {
	pid := os.Getpid()
	p, err := GetProcess(pid, nil)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, pid, p.PID)
}
