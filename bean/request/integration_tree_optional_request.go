package request

type TreeNodeSort struct {
	//enum (specialty, systemType, floor, category, family, familyType, building, unit, roomType, room, subFamilyType)
	NodeType int `json:"nodeType"` //样例 : "3"

	SortBy       string   `json:"sortBy"`
	SortedValues []string `json:"sortedValues"`
}

type IntegrationTreeOptionalRequest struct {
	CustomizedNodeKeys   map[string]string           `json:"customizedNodeKeys,omitempty"`
	FileIdElementIds     []*FileIdHashWithElementIds `json:"fileIdElementIds,omitempty"`
	Sorts                []*TreeNodeSort             `json:"Sorts,omitempty"`
	SortedNamesHierarchy [][]string                  `json:"sortedNamesHierarchy,omitempty"`
}

func NewIntegrationTreeOptionalRequest() *IntegrationTreeOptionalRequest {
	o := &IntegrationTreeOptionalRequest{}
	return o
}
