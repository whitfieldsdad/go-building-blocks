package bb

import (
	"github.com/shirou/gopsutil/v3/process"
)

type ProcessTree struct {
	pidToPpid map[int]int
}

func NewProcessTree() *ProcessTree {
	return &ProcessTree{
		pidToPpid: make(map[int]int),
	}
}

func GetProcessTree() (*ProcessTree, error) {
	t := NewProcessTree()
	ps, err := process.Processes()
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		pid := int(p.Pid)
		ppid32, err := p.Ppid()
		if err != nil {
			continue
		}
		ppid := int(ppid32)
		t.AddProcess(ppid, pid)
	}
	return t, nil
}

func (t *ProcessTree) AddProcess(ppid, pid int) error {
	t.pidToPpid[pid] = ppid
	return nil
}

func (t *ProcessTree) RemoveProcesses(pids ...int) {
	for _, pid := range pids {
		delete(t.pidToPpid, pid)
	}
}

func (t ProcessTree) GetAncestorPids(pid int) []int {
	ancestors := []int{}
	for {
		ppid, ok := t.pidToPpid[pid]
		if !ok {
			break
		}
		ancestors = append(ancestors, ppid)
		pid = ppid
	}
	return ancestors
}

func (t ProcessTree) GetDescendantPids(pid int) []int {
	var descendants []int
	children := t.GetChildPids(pid)
	for _, child := range children {
		descendants = append(descendants, child)
		descendants = append(descendants, t.GetDescendantPids(child)...)
	}
	return descendants
}

func (t ProcessTree) GetParentPid(pid int) (int, bool) {
	ppid, ok := t.pidToPpid[pid]
	return ppid, ok
}

func (t ProcessTree) GetChildPids(pid int) []int {
	children := []int{}
	for p, parent := range t.pidToPpid {
		if pid == parent {
			children = append(children, p)
		}
	}
	return children
}

func (t ProcessTree) GetSiblingPids(pid int) []int {
	ppid, ok := t.GetParentPid(pid)
	if !ok {
		return nil
	}
	var siblings []int
	for _, child := range t.GetChildPids(ppid) {
		if child != pid {
			siblings = append(siblings, child)
		}
	}
	return siblings
}
