// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

import "github.com/chnykn/bimface/bean/common"

type IntegrateSource struct {
	FileId   int64 `json:"fileId"` //必填
	FileName int64 `json:"fileName"`

	Building  string `json:"building"`
	DatabagId string `json:"databagId"`

	Specialty     string  `json:"specialty"`
	SpecialtySort float64 `json:"specialtySort,omitempty"`

	Floor     string  `json:"floor,omitempty"`
	FloorSort float64 `json:"floorSort,omitempty"`

	Transform []float64 `json:"transform,omitempty"`
}

func NewIntegrateSource(fileId int64, floor, specialty string) *IntegrateSource {
	o := &IntegrateSource{
		FileId:    fileId,
		Floor:     floor,
		Specialty: specialty,
	}
	return o
}

//---------------------------------------------------------------------

type FloorMappingItem struct {
	FileFloorId      string `json:"fileFloorId"`
	ProjectFloorId   string `json:"projectFloorId"`
	ProjectFloorName string `json:"projectFloorName"`
}

//---------------------------------------------------------------------

type FileIntegrateRequest struct {
	Name              string `json:"name"`
	ParentIntegrateId int64  `json:"parentIntegrateId"`

	SourceId string             `json:"sourceId,omitempty"`
	Sources  []*IntegrateSource `json:"sources,omitempty"`

	FloorMapping []*FloorMappingItem `json:"floorMapping,omitempty"`

	FloorSort     []string `json:"floorSort,omitempty"`
	SpecialtySort []string `json:"specialtySort,omitempty"`

	PropertyOverrides []*common.PropertyOverride `json:"propertyOverrides,omitempty"`

	RuleFileIds []interface{} `json:"ruleFileIds,omitempty"`

	Config            map[string]string `json:"config,omitempty"`
	InternalConfigMap map[string]string `json:"internalConfigMap,omitempty"`

	Callback string `json:"callback,omitempty"`
}

func NewFileIntegrateRequest(name string) *FileIntegrateRequest {
	o := &FileIntegrateRequest{
		Name:     name,
		Callback: "",
		Sources:  make([]*IntegrateSource, 0),
	}
	return o
}
