// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/bean/common"

type SplitDrawingFrame struct {
	Id          int64               `json:"FileId"`
	Name        string              `json:"name"`
	Number      string              `json:"number"`
	BoundingBox *common.BoundingBox `json:"boundingBox"`
}

type SplitDrawingFrames struct {
	Id     int64                `json:"FileId"`
	Name   string               `json:"name"`
	Frames []*SplitDrawingFrame `json:"frames"`
}
