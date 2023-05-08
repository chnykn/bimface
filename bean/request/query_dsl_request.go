package request

type BoolOp struct {
	Contain map[string]any `json:"contain,omitempty"`
	Match   map[string]any `json:"match,omitempty"`
	BoolAnd []BoolOp       `json:"boolAnd,omitempty"`
	BoolOr  []BoolOp       `json:"boolOr,omitempty"`
}

type QueryDSLRequest struct {
	TargetType string   `json:"targetType,omitempty"`
	TargetIds  []string `json:"targetIds,omitempty"`
	Query      BoolOp   `json:"query,omitempty"`
}
