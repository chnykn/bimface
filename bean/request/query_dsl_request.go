package request

type DSLCondition struct {
	Contain map[string]any `json:"contain,omitempty"`
	Match   map[string]any `json:"match,omitempty"`
	BoolAnd []*DSLCondition `json:"boolAnd,omitempty"`
	BoolOr  []*DSLCondition `json:"boolOr,omitempty"`
}

type QueryDSLRequest struct {
	TargetType string        `json:"targetType,omitempty"`
	TargetIds  []string      `json:"targetIds,omitempty"`
	Query      *DSLCondition `json:"query,omitempty"`
}
