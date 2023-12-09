package bb

type Host struct {
	Id                string             `json:"id"`
	CPUs              []CPU              `json:"cpus,omitempty"`
	Motherboard       Motherboard        `json:"motherboard,omitempty"`
	MemoryModules     []MemoryModule     `json:"memory_modules,omitempty"`
	NetworkInterfaces []NetworkInterface `json:"network_interfaces,omitempty"`
	NetworkLocation   NetworkLocation    `json:"network_location,omitempty"`
	OperatingSystem   OperatingSystem    `json:"operating_system,omitempty"`
}

func GetHost() (*Host, error) {
	id, err := GetHostId()
	if err != nil {
		return nil, err
	}
	os := GetOperatingSystem()
	networkLocation, _ := GetNetworkLocation()
	networkInterfaces, _ := ListNetworkInterfaces()
	motherboard, _ := GetMotherboard()
	cpus, _ := ListCPUs()

	host := &Host{
		Id:                id,
		CPUs:              cpus,
		Motherboard:       *motherboard,
		NetworkInterfaces: networkInterfaces,
		NetworkLocation:   *networkLocation,
		OperatingSystem:   *os,
	}
	return host, nil
}
