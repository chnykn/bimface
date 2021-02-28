// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

//TreeNode ***
type TreeNode struct {
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	Name         string      `json:"name"`
	ActualName   string      `json:"actualName"`
	Data         interface{} `json:"data,omitempty"`
	ElementCount int64       `json:"elementCount"`
	Items        []TreeNode
}

//NewTreeNode ***
func NewTreeNode() *TreeNode {
	o := &TreeNode{
		// Type: typ
		// FileId:   id,
		// Name: name,
		//ElementCount: 0,
		Items: make([]TreeNode, 0),
	}
	return o
}

//--------------------------------

type Tree struct {
	Root  string      `json:"root"`
	Items []*TreeNode `json:"items"`
}
