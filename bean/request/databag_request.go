// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

//DataBagRequest ***
type DataBagRequest struct {
	FileId         *int64 `json:"fileId,omitempty"` // FileId IntegrateId CompareId 三者只能有一个填写
	IntegrateId    *int64 `json:"integrateId,omitempty"`
	CompareId      *int64 `json:"compareId,omitempty"`
	Callback       string `json:"callback,omitempty"`
	DataBagVersion string `json:"databagVersion,omitempty"`
}

func NewDataBagRequest() *DataBagRequest {
	//fileId string, integrateId string, compareId string,	callback string, databagVersion string

	o := &DataBagRequest{
		// FileId:         fileId,
		// IntegrateId:    integrateId,
		// CompareId:      compareId,
		// Callback:       callback,
		// DataBagVersion: databagVersion,
	}

	return o
}
