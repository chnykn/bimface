// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//CategoryNode ***
type CategoryNode struct {
	ActualName   string
	Type         string
	ID           string
	Name         string
	Data         string
	ElementCount int64
	Items        []CategoryNode
}

//NewCategoryNode ***
func NewCategoryNode() *CategoryNode { //typ string, id string, name string
	o := &CategoryNode{
		// Type: typ
		// ID:   id,
		// Name: name,
		ElementCount: 0,
		Items:        make([]CategoryNode, 0),
	}
	return o
}
