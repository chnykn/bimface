// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type IntegrateFileData struct {
	IntegrateId int64 `json:"integrateId"`

	FileId   int64  `json:"fileId"`
	FileName string `json:"fileName"`
	LinkedBy string `json:"linkedBy"`

	DatabagId         string `json:"databagId"`
	DrawingSheetCount int    `json:"drawingSheetCount"`

	Floor     string  `json:"floor"`
	FloorSort float64 `json:"floorSort"`

	Specialty     string  `json:"specialty"`
	SpecialtySort float64 `json:"specialtySort"`
}

type IntegrateFileBean IntegrateFileData
