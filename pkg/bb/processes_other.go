//go:build !windows
// +build !windows

package bb

import "os"

func IsElevated() (bool, error) {
	return os.Geteuid() == 0, nil
}
