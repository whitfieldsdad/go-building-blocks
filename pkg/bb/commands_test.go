package bb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteNativeCommand(t *testing.T) {
	result, err := ExecuteCommand(context.Background(), "whoami", Native, nil)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, result.ExitCode)
	assert.Equal(t, "whoami", result.Command.Command)
	assert.Equal(t, DefaultCommandType, result.Command.Type)

	// Ensure that we only look up the subprocess by default.
	assert.Equal(t, 1, len(result.GetProcesses()))
}

func TestExecuteCommandAndIncludeParentProcesses(t *testing.T) {
	opts := &CommandOptions{
		IncludeParentProcesses: true,
	}
	result, err := ExecuteCommand(context.Background(), "whoami", string(DefaultCommandType), opts)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Greater(t, len(result.GetProcesses()), 1)
}
