package response

/**
  "code" : "perimeter",
  "extension" : "object",
  "key" : "perimeter",
  "orderNumber" : 0,
  "unit" : "mm",
  "value" : 17200,
  "valueType" : 2
*/
type MaterialItem struct {
	Code        string `json:"code"`
	Extension   string `json:"extension"`
	Key         string `json:"key"`
	OrderNumber int    `json:"orderNumber"`
	Unit        string `json:"unit"`
	Value       int    `json:"value"`
	ValueType   int    `json:"valueType"`
}

type MaterialParam struct {
	Group string          `json:"group"`
	Items []*MaterialItem `json:"items"`
}

type ElementMaterial struct {
	Id         string           `json:"id"`
	Name       string           `json:"name"`
	Parameters []*MaterialParam `json:"parameters"`
}
