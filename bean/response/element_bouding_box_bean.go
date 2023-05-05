// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/v3/bean/common"

type ElementIdWithBoundingBox struct {
	ElementId   string             `json:"elementId"`
	BoundingBox common.BoundingBox `json:"boundingBox"`
}

type ElementBoundingBoxBean ElementIdWithBoundingBox
