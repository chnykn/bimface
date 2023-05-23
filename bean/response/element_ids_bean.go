package response

type ElementIds struct {
	TargetId   string   `json:"targetId"`
	ElementIds []string `json:"elementIds"`
}

type ElementIdsBean ElementIds

type ElementIdsArrayBean []*ElementIds
