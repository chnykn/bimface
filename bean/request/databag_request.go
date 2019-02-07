// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

//DatabagRequest ***
type DatabagRequest struct {
	FileID         *int64 `json:"fileId,omitempty"` // FileID IntegrateID CompareID 三者只能有一个填写
	IntegrateID    *int64 `json:"integrateId,omitempty"`
	CompareID      *int64 `json:"compareId,omitempty"`
	Callback       string `json:"callback,omitempty"`
	DatabagVersion string `json:"databagVersion,omitempty"`
}

//NewDatabagRequest ***
func NewDatabagRequest() *DatabagRequest {
	//fileID string, integrateID string, compareID string,	callback string, databagVersion string

	o := &DatabagRequest{
		// FileID:         fileID,
		// IntegrateID:    integrateID,
		// CompareID:      compareID,
		// Callback:       callback,
		// DatabagVersion: databagVersion,
	}

	return o
}
