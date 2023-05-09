package response

type PropertyValuesBean struct {
	Property string   `json:"property"`
	Values   []string `json:"values"`
}

type PropertyValuesBeanList []*PropertyValuesBean
