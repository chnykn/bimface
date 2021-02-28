package response

import "github.com/chnykn/bimface/v2/bean/common"

type Area struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Boundary  string  `json:"boundary"`
	ViewName  string  `json:"viewName"`
	LevelId   string  `json:"levelId"`
	Area      float64 `json:"area"`
	Perimeter float64 `json:"perimeter"`

	MaxPt *common.Coordinate `json:"maxPt,omitempty"`
	MinPt *common.Coordinate `json:"minPt,omitempty"`

	Properties []*common.PropertyGroup `json:"properties"`
}

type AreaBean Area
