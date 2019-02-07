// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

//"Intgr" stand for "Integration"

//IntgrSource ***
type IntgrSource struct {
	FileID        int64   `json:"fileId"` //必填
	Specialty     string  `json:"specialty"`
	SpecialtySort float64 `json:"specialtySort,omitempty"`
	Floor         string  `json:"floor,omitempty"`
	FloorSort     float64 `json:"floorSort,omitempty"`
}

//NewIntgrSource ***
func NewIntgrSource(fileID int64, specialty string) *IntgrSource {
	o := &IntgrSource{
		FileID:    fileID,
		Specialty: specialty,
	}
	return o
}

//---------------------------------------------------------------------

const (
	defualtIntegratePriority byte = 2
)

//IntgrRequest ***
type IntgrRequest struct {
	Name     string             `json:"name"`     // 必填
	Priority byte               `json:"priority"` // [1,2,3]  数字越大，优先级越低
	Callback string             `json:"callback,omitempty"`
	SourceID string             `json:"sourceID,omitempty"`
	Sources  []*IntgrSource `json:"sources"` //必填
}

//NewIntgrRequest ***
func NewIntgrRequest(name string) *IntgrRequest { //, sourceID string
	o := &IntgrRequest{
		Name:     name,
		Priority: defualtIntegratePriority,
		Callback: "",
		//SourceID: sourceID,
		Sources: make([]*IntgrSource, 0),
	}
	return o
}
