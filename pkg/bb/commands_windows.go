package bb

import "syscall"

func getSysProcAttrs() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		HideWindow: true,
	}
}
