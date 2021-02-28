package response

type LinkGraphNode struct {
	Name      string `json:"name"`
	FileId    string `json:"fileId"`
	DatabagId string `json:"databagId"`

	LinkName      string `json:"linkName"`
	LinkPathHash  string `json:"linkPathHash"`
	LinkTransform string `json:"linkTransform"`

	Links  []*LinkGraphNode `json:"links,omitempty"`
	Params []interface{}    `json:"params,omitempty"`
}

type LinkGraphNodeBean LinkGraphNode
