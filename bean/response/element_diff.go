// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//ElementAttr ***
type ElementAttr struct {
	Key   string `json:"key"`            //属性名
	Value string `json:"value"`          //属性值
	Unit  string `json:"unit,omitempty"` //单位
}

//ChangeAttrs ***
type ChangeAttrs struct {
	A ElementAttr
	B ElementAttr
}

// ElementQty ***
type ElementQty struct {
	Name string  `json:"name"` //工程量名称(用于显示)
	Code string  `json:"code"` //工程量编码(用于运算)
	Desc string  `json:"des"`  //工程量描述
	Unit string  `json:"unit"` //单位
	Qty  float64 `json:"qty"`  //数值
}

//ChangeQtys ***
type ChangeQtys struct {
	A ElementQty
	B ElementQty
}

//ElementDiff ***
type ElementDiff struct {
	A string //变化图元前一个版本的ID
	B string //变化图元后一个版本的ID

	NewAttributes    []*ElementAttr `json:"newAttributes"`
	DeleteAttributes []*ElementAttr `json:"deleteAttributes"`
	ChangeAttributes []*ChangeAttrs `json:"changeAttributes"`

	NewQuantities    []*ElementQty `json:"newQuantities"`
	DeleteQuantities []*ElementQty `json:"deleteQuantities"`
	ChangeQuantities []*ChangeQtys `json:"changeQuantities"`
}

//NewElementDiff ***
func NewElementDiff() *ElementDiff {
	o := &ElementDiff{
		NewAttributes:    make([]*ElementAttr, 0),
		DeleteAttributes: make([]*ElementAttr, 0),
		ChangeAttributes: make([]*ChangeAttrs, 0),

		NewQuantities:    make([]*ElementQty, 0),
		DeleteQuantities: make([]*ElementQty, 0),
		ChangeQuantities: make([]*ChangeQtys, 0),
	}
	return o
}
