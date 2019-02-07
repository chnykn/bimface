// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

//IntegrateSource ***
type IntegrateSource struct {
	FileID        int64   `json:"fileId"` //必填
	Specialty     string  `json:"specialty"`
	SpecialtySort float64 `json:"specialtySort,omitempty"`
	Floor         string  `json:"floor,omitempty"`
	FloorSort     float64 `json:"floorSort,omitempty"`
}

//NewIntegrateSource ***
func NewIntegrateSource(fileID int64, specialty string) *IntegrateSource {
	o := &IntegrateSource{
		FileID:    fileID,
		Specialty: specialty,
	}
	return o
}

//---------------------------------------------------------------------

const (
	defualtIntegratePriority byte = 2
)

//IntegrateRequest ***
type IntegrateRequest struct {
	Name     string             `json:"name"`     // 必填
	Priority byte               `json:"priority"` // [1,2,3]  数字越大，优先级越低
	Callback string             `json:"callback,omitempty"`
	SourceID string             `json:"sourceID,omitempty"`
	Sources  []*IntegrateSource `json:"sources"` //必填
}

//NewIntegrateRequest ***
func NewIntegrateRequest(name string) *IntegrateRequest { //, sourceID string
	o := &IntegrateRequest{
		Name:     name,
		Priority: defualtIntegratePriority,
		Callback: "",
		//SourceID: sourceID,
		Sources: make([]*IntegrateSource, 0),
	}
	return o
}
