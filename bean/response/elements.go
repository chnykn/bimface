// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/bean/common"

//Element ***
type Element struct {
	FileID    string `json:"fileId"`
	ElementID string `json:"elementId"`
}

//NewElement ***
func NewElement(fileID string, elementID string) *Element {
	o := &Element{
		FileID:    fileID,
		ElementID: elementID,
	}
	return o
}

//---------------------------------------------------------------------

//Elements ***
type Elements struct {
	Elements    []Element          `json:"elements"`
	BoundingBox common.BoundingBox `json:"boundingBox"`
}

//NewElements ***
func NewElements() *Elements { //id int64, name string
	o := &Elements{
		Elements: make([]Element, 0),
		//BoundingBox: common.NewBoundingBox(),
	}
	return o
}
