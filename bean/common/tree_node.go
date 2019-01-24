// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

//TreeNode ***
type TreeNode struct {
	Type         string
	ID           string
	Name         string
	ElementCount int64
	Items        []TreeNode
}

//NewTreeNode ***
func NewTreeNode() *TreeNode { //typ string, id string, name string
	o := &TreeNode{
		// Type: typ
		// ID:   id,
		// Name: name,
		ElementCount: 0,
		Items:        make([]TreeNode, 0),
	}
	return o
}
