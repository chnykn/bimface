// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/bean/common"

type ModelCompareChange struct {
	A string `json:"_A"`
	B string `json:"_B"`

	NewAttributes []*common.Attribute `json:"newAttributes"`
	NewQuantities []*common.Quantity  `json:"newQuantities"`

	DeleteAttributes []*common.Attribute `json:"deleteAttributes"`
	DeleteQuantities []*common.Quantity  `json:"deleteQuantities"`

	ChangeAttributes []*common.ChangeAttribute `json:"changeAttributes"`
	ChangeQuantities []*common.ChangeQuantity  `json:"changeQuantities"`
}

type ModelCompareChangeBean ModelCompareChange
