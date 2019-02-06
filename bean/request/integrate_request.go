// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

//IntegrateSource ***
type IntegrateSource struct {
	FileID        int64 //必填
	Specialty     string
	SpecialtySort float64
	Floor         string
	FloorSort     float64
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
	Name     string // 必填
	Priority byte   // [1,2,3]  数字越大，优先级越低
	Callback string
	SourceID string             //TODO: shoud be *string type
	Sources  []*IntegrateSource //必填
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
