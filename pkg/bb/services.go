package bb

type Service struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name,omitempty"`
	Description string   `json:"description,omitempty"`
	State       string   `json:"state,omitempty"`
	Process     *Process `json:"process,omitempty"`
}

func ListServices(opts *ProcessOptions) ([]Service, error) {
	return listServices(opts)
}
