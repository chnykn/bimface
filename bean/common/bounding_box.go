// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

//BoundingBox ***
type BoundingBox struct {
	Min Coordinate `json:"min"`
	Max Coordinate `json:"max"`
}

//NewBoundingBox ***
func NewBoundingBox() *BoundingBox {
	o := &BoundingBox{}
	return o
}
