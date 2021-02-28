package response

type NetworkNode struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type MEPSystem struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	SystemType     string `json:"systemType"`
	SystemCategory string `json:"systemCategory"`
	BaseEquipment  string `json:"baseEquipment"`

	Terminals []string       `json:"terminals,omitempty"`
	Network   []*NetworkNode `json:"network,omitempty"`
}

type MEPSysBean MEPSystem
