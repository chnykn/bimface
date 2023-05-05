// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

type Quantity struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
	Unit string `json:"unit"`
	Desc string `json:"decs,omitempty"`
}

type ChangeQuantity struct {
	A *Quantity `json:"_A"`
	B *Quantity `json:"_B"`
}
