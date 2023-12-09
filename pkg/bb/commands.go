package bb

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type CommandOptions struct {
	IncludeAll             bool            `json:"include_all,omitempty"`
	IncludeOutput          bool            `json:"include_output,omitempty"`
	IncludeParentProcesses bool            `json:"include_parent_processes,omitempty"`
	ProcessOptions         *ProcessOptions `json:"process_options,omitempty"`
}

func NewCommandOptions() *CommandOptions {
	return &CommandOptions{
		IncludeAll: true,
	}
}

type Command struct {
	Command        string                 `json:"command"`
	Type           string                 `json:"type"`
	InputArguments map[string]interface{} `json:"input_arguments"`
}

func NewCommand(command, commandType string) *Command {
	return &Command{
		Command: command,
		Type:    commandType,
	}
}

func (command Command) Execute(ctx context.Context, opts *CommandOptions) (*ExecutedCommand, error) {
	return ExecuteCommand(ctx, command.Command, command.Type, opts)
}

type ExecutedCommand struct {
	Id          string    `json:"id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Command     Command   `json:"command"`
	ExitCode    int       `json:"exit_code"`
	Process     Process   `json:"subprocess"`
	ProcessTree []Process `json:"process_tree"`
}

func (result ExecutedCommand) GetProcesses() []Process {
	var processes []Process
	processes = append(processes, result.Process)
	processes = append(processes, result.ProcessTree...)
	return processes
}

func (result ExecutedCommand) GetDuration() time.Duration {
	return result.EndTime.Sub(result.StartTime)
}

func ExecuteCommand(ctx context.Context, command, commandType string, opts *CommandOptions) (*ExecutedCommand, error) {
	if opts == nil {
		opts = NewCommandOptions()
	}
	argv, err := wrapCommand(command, commandType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to wrap command")
	}
	startTime := time.Now()
	process, err := executeArgv(ctx, argv, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute command")
	}
	endTime := time.Now()

	var tree []Process
	if opts.IncludeParentProcesses {
		pid := os.Getpid()
		tree, _ = GetProcessAncestors(pid)
		self, err := GetProcess(pid, opts.ProcessOptions)
		if err == nil {
			tree = append(tree, *self)
		}
	}
	executedCommand := &ExecutedCommand{
		Id:          NewUUID4(),
		StartTime:   startTime,
		EndTime:     endTime,
		Command:     Command{Command: command, Type: commandType},
		ExitCode:    *process.ExitCode,
		Process:     *process,
		ProcessTree: tree,
	}
	return executedCommand, nil
}

func executeArgv(ctx context.Context, argv []string, opts *CommandOptions) (*Process, error) {
	path, err := exec.LookPath(argv[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to find command")
	}
	cmd := exec.Command(path, argv[1:]...)
	cmd.SysProcAttr = getSysProcAttrs()

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	// Execute the command.
	err = cmd.Start()
	if err != nil {
		return nil, errors.Wrap(err, "failed to start command")
	}

	// Collect information about the subprocess.
	pid := cmd.Process.Pid
	process, err := GetProcess(pid, opts.ProcessOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to collect process metadata")
	}
	if process.Argv == nil || process.CommandLine == "" {
		process.Argv = argv
		process.CommandLine = strings.Join(argv, " ")
	}

	// Wait for the command to complete.
	err = cmd.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "failed to wait for command to exit")
	}
	process.Stdout = stdout.String()
	process.Stderr = stderr.String()

	exitCode := cmd.ProcessState.ExitCode()
	process.ExitCode = &exitCode
	return process, nil
}

var (
	WindowsPowerShell = "powershell"
	PowerShellCore    = "pwsh"
	PowerShell        = getPowerShellCommandType()
	CommandPrompt     = "command_prompt"
	Sh                = "sh"
	Bash              = "bash"
	Native            = "native"
)

var (
	DefaultCommandType = Native
)

var (
	commandShims = map[string][]string{
		WindowsPowerShell: {"powershell", "-ExecutionPolicy", "Bypass", "-Command"},
		PowerShellCore:    {"pwsh", "-Command"},
		CommandPrompt:     {"cmd", "/c"},
		Sh:                {"sh", "-c"},
		Bash:              {"bash", "-c"},
	}
)

func getPowerShellCommandType() string {
	if runtime.GOOS == "windows" {
		return WindowsPowerShell
	}
	return PowerShellCore
}

func wrapCommand(command, commandType string) ([]string, error) {
	if commandType == "" {
		commandType = DefaultCommandType
	}
	if commandType == Native {
		return []string{command}, nil
	}
	argv := commandShims[commandType]
	if argv == nil {
		return nil, errors.Errorf("invalid command type: %s", commandType)
	}
	argv = append(argv, command)
	return argv, nil
}
