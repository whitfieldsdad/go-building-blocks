package bb

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOperatingSystem(t *testing.T) {
	os := GetOperatingSystem()
	assert.NotNil(t, os, "Operating system should not be nil")
	assert.Equal(t, runtime.GOOS, os.Type, "OS type should match")
	assert.Equal(t, runtime.GOARCH, os.Arch, "OS architecture should match")
}
