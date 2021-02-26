package response

import "github.com/chnykn/bimface/bean/common"

type Area struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	ViewName  string  `json:"viewName"`
	Boundary  string  `json:"boundary"`
	LevelId   string  `json:"levelId"`
	Area      float64 `json:"area"`
	Perimeter float64 `json:"perimeter"`

	MaxPt *common.Coordinate `json:"maxPt,omitempty"`
	MinPt *common.Coordinate `json:"minPt,omitempty"`

	BBoxMax *common.Coordinate `json:"bboxMax,omitempty"`
	BBoxMin *common.Coordinate `json:"bboxMin,omitempty"`

	Properties []*PropertyGroup `json:"properties"`
}
