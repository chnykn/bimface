// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/v3/bean/common"

type ObjectOnFloor struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Boundary string `json:"boundary"`
	LevelId  string `json:"levelId"`

	MaxPt *common.Coordinate `json:"maxPt,omitempty"`
	MinPt *common.Coordinate `json:"minPt,omitempty"`
}

type Floor struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Elevation  float64 `json:"elevation"`
	ArchElev   float64 `json:"archElev"`
	StructElev float64 `json:"structElev"`
	Height     float64 `json:"height"`
	MiniMap    string  `json:"miniMap"`

	Areas []*ObjectOnFloor `json:"areas,omitempty"`
	Rooms []*ObjectOnFloor `json:"rooms,omitempty"`
}

//-----------------------------------------

type FloorBean Floor

type FloorsBean struct {
	FileId int64    `json:"FileId"`
	Floors []*Floor `json:"floors,omitempty"`
}
