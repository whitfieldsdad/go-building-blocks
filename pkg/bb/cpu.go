package bb

import (
	"github.com/jaypipes/ghw"
)

type CPU struct {
	Name         string   `json:"name"`
	Model        string   `json:"model"`
	Vendor       string   `json:"vendor"`
	TotalCores   int      `json:"total_cores"`
	TotalThreads int      `json:"total_threads"`
	Capabilities []string `json:"capabilities,omitempty"`
}

func ListCPUs() ([]CPU, error) {
	cpu, err := ghw.CPU()
	if err != nil {
		return nil, err
	}
	var cpus []CPU
	for _, info := range cpu.Processors {
		cpu := CPU{
			Name:         info.Model,
			Model:        info.Model,
			Vendor:       info.Vendor,
			TotalCores:   int(info.NumCores),
			TotalThreads: int(info.NumThreads),
			Capabilities: info.Capabilities,
		}
		cpus = append(cpus, cpu)
	}
	return cpus, nil
}
