// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

//TreeNode ***
type TreeNode struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	ActualName   string `json:"actualName"`
	Data         string `json:"data,omitempty"`
	ElementCount int64  `json:"elementCount"`
	Items        []TreeNode
}

//NewTreeNode ***
func NewTreeNode() *TreeNode {
	o := &TreeNode{
		// Type: typ
		// ID:   id,
		// Name: name,
		//ElementCount: 0,
		Items: make([]TreeNode, 0),
	}
	return o
}
