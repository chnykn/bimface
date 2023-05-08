package request

type Condition struct {
	Contain map[string]any `json:"contain,omitempty"`
	Match   map[string]any `json:"match,omitempty"`
	BoolAnd []Condition    `json:"boolAnd,omitempty"`
	BoolOr  []Condition    `json:"boolOr,omitempty"`
}

type QueryRequest struct {
	TargetType string    `json:"targetType,omitempty"`
	TargetIds  []string  `json:"targetIds,omitempty"`
	Query      Condition `json:"query,omitempty"`
}
