// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/bean/common"

//"Intgr" stand for "Integration"

//IntgrElement ***
type IntgrElement struct {
	FileId    string `json:"fileId"`
	ElementId string `json:"elementId"`
}

//IntgrElements ***
type IntgrElements struct {
	Elements    []*IntgrElement    `json:"elements"`
	BoundingBox common.BoundingBox `json:"boundingBox"`
}
