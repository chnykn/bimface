// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

type BoundingBox struct {
	Min *Coordinate `json:"min"`
	Max *Coordinate `json:"max"`
}

func NewBoundingBox() *BoundingBox {
	o := &BoundingBox{}
	return o
}
