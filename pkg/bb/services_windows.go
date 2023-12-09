package bb

import (
	"github.com/yusufpapurcu/wmi"
)

type Win32_Service struct {
	Name        string
	DisplayName string
	Description string
	ProcessId   uint32
	State       string
}

func listServices(opts *ProcessOptions) ([]Service, error) {
	if opts == nil {
		opts = NewProcessOptions()
	}
	var win32Services []Win32_Service
	err := wmi.Query("SELECT * FROM Win32_Service", &win32Services)
	if err != nil {
		return nil, err
	}
	var svcs []Service
	for _, s := range win32Services {
		svc := Service{
			Name:        s.Name,
			DisplayName: s.DisplayName,
			Description: s.Description,
			State:       s.State,
		}
		if s.State == "Running" {
			svc.Process, err = GetProcess(int(s.ProcessId), opts)
			if err != nil {
				return nil, err
			}
		}
		svcs = append(svcs, svc)
	}
	return svcs, nil
}
