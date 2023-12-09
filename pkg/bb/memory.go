package bb

import "github.com/jaypipes/ghw"

type MemoryModule struct {
	Label        string `json:"label"`
	Location     string `json:"location"`
	SerialNumber string `json:"serial_number"`
	Size         int64  `json:"size_bytes"`
	Vendor       string `json:"vendor"`
}

func ListMemoryModules() ([]MemoryModule, error) {
	memory, err := ghw.Memory()
	if err != nil {
		return nil, err
	}
	var rows []MemoryModule
	for _, info := range memory.Modules {
		row := MemoryModule{
			Label:        info.Label,
			Location:     info.Location,
			SerialNumber: info.SerialNumber,
			Size:         info.SizeBytes,
			Vendor:       info.Vendor,
		}
		rows = append(rows, row)
	}
	return rows, nil
}
