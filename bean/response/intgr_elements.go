// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/bean/common"

//"Intgr" stand for "Integration"

//IntgrElement ***
type IntgrElement struct {
	FileID    string `json:"fileId"`
	ElementID string `json:"elementId"`
}

//NewIntgrElement ***
func NewIntgrElement(fileID string, elementID string) *IntgrElement {
	o := &IntgrElement{
		FileID:    fileID,
		ElementID: elementID,
	}
	return o
}

//---------------------------------------------------------------------

//IntgrElements ***
type IntgrElements struct {
	Elements    []*IntgrElement    `json:"elements"`
	BoundingBox common.BoundingBox `json:"boundingBox"`
}

//NewElements ***
func NewElements() *IntgrElements {
	o := &IntgrElements{
		Elements: make([]*IntgrElement, 0),
		//BoundingBox: common.NewBoundingBox(),
	}
	return o
}
