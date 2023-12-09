package bb

import (
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type Process struct {
	Id          string     `json:"id" yaml:"id"`
	Time        time.Time  `json:"time" yaml:"time"`
	Name        string     `json:"name,omitempty"`
	PID         int        `json:"pid"`
	PPID        *int       `json:"ppid,omitempty"`
	User        *User      `json:"user,omitempty"`
	Executable  *File      `json:"executable,omitempty"`
	CommandLine string     `json:"command_line,omitempty"`
	Argv        []string   `json:"argv,omitempty"`
	Argc        int        `json:"argc,omitempty"`
	CreateTime  *time.Time `json:"create_time,omitempty"`
	ExitTime    *time.Time `json:"exit_time,omitempty"`
	ExitCode    *int       `json:"exit_code,omitempty"`
	Stdout      string
	Stderr      string
}

func GetProcess(pid int) (*Process, error) {
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		return nil, err
	}
	ppid32, err := p.Ppid()
	if err != nil {
		return nil, err
	}
	ppid := int(ppid32)

	var file *File
	exe, _ := p.Exe()
	if exe != "" {
		file, _ = GetFile(exe)
	}
	name, _ := p.Name()
	argv, _ := p.CmdlineSlice()
	argc := len(argv)
	cmdline := strings.Join(argv, " ")

	var (
		startTime *time.Time
	)
	startTimeMs, err := p.CreateTime()
	if err == nil {
		startTime = ParseUnixTimestamp(startTimeMs)
	}
	return &Process{
		Name:        name,
		PID:         pid,
		PPID:        &ppid,
		Executable:  file,
		Argv:        argv,
		Argc:        argc,
		CommandLine: cmdline,
		CreateTime:  startTime,
	}, nil
}

func GetProcessAncestors(pid int) ([]Process, error) {
	panic("not implemented")
}

func GetProcessDescendants(pid int) ([]Process, error) {
	panic("not implemented")
}

func GetProcessSiblings(pid int) ([]Process, error) {
	panic("not implemented")
}
