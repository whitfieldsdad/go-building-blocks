package bb

import (
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcessOptions struct {
	IncludeAll          bool         `json:"include_all,omitempty"`
	IncludeFileMetadata bool         `json:"include_file_metadata,omitempty"`
	IncludeFileHashes   bool         `json:"include_file_hashes,omitempty"`
	FileOptions         *FileOptions `json:"file_options,omitempty"`
}

func NewProcessOptions() *ProcessOptions {
	return &ProcessOptions{
		IncludeAll:  true,
		FileOptions: NewFileOptions(),
	}
}

type Process struct {
	Id          string     `json:"id"`
	Time        time.Time  `json:"time"`
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
	Stdout      string     `json:"stdout,omitempty"`
	Stderr      string     `json:"stderr,omitempty"`
}

func GetProcess(pid int, opts *ProcessOptions) (*Process, error) {
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		return nil, err
	}
	ppid32, err := p.Ppid()
	if err != nil {
		return nil, err
	}
	ppid := int(ppid32)

	if opts == nil {
		opts = NewProcessOptions()
	}

	var file *File
	if opts.IncludeAll || opts.IncludeFileMetadata {
		exe, _ := p.Exe()
		if exe != "" {
			file, _ = GetFile(exe, opts.FileOptions)
		}
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
	var user *User
	username, err := p.Username()
	if err == nil {
		user, _ = GetUser(username)
	}
	id, err := GetProcessUUID(pid, ppid)
	if err != nil {
		return nil, err
	}
	return &Process{
		Id:          id,
		Time:        time.Now(),
		Name:        name,
		PID:         pid,
		PPID:        &ppid,
		User:        user,
		Executable:  file,
		Argv:        argv,
		Argc:        argc,
		CommandLine: cmdline,
		CreateTime:  startTime,
	}, nil
}

func GetProcessUUID(pid, ppid int) (string, error) {
	m := map[string]interface{}{
		"pid":  pid,
		"ppid": ppid,
	}
	return NewUUID5FromMap(DefaultUUIDNamespace, m)
}

func ListProcesses(opts *ProcessOptions) ([]Process, error) {
	pids, err := process.Pids()
	if err != nil {
		return nil, err
	}
	processes := make([]Process, len(pids))
	for i, pid := range pids {
		p, err := GetProcess(int(pid), opts)
		if err != nil {
			continue
		}
		processes[i] = *p
	}
	return processes, nil
}
