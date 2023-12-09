package bb

import "github.com/jaypipes/ghw"

type Motherboard struct {
	Product      string `json:"product"`
	AssetTag     string `json:"asset_tag"`
	SerialNumber string `json:"serial_number"`
	Vendor       string `json:"vendor"`
	Version      string `json:"version"`
}

func GetMotherboard() (*Motherboard, error) {
	info, err := ghw.Baseboard()
	if err != nil {
		return nil, err
	}
	motherboard := Motherboard{
		Product:      info.Product,
		AssetTag:     info.AssetTag,
		SerialNumber: info.SerialNumber,
		Vendor:       info.Vendor,
		Version:      info.Version,
	}
	return &motherboard, nil
}
