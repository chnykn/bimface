// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

// TreeNode ***
type TreeNode struct {
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	Name         string      `json:"name"`
	ActualName   string      `json:"actualName"`
	Data         interface{} `json:"data,omitempty"`
	ElementCount int64       `json:"elementCount"`
	Items        []*TreeNode `json:"items"`

	// 为兼容离线数据包中 data\tree.json 文件内 familyType 节点下element列表
	FileId     string   `json:"fileId"`
	ElementIds []string `json:"elementIds"`

	//---- 额外增加，其他用途 -----
	Parent *TreeNode   `json:"-"`
	ExData interface{} `json:"-"`
}

// NewTreeNode ***
func NewTreeNode() *TreeNode {
	o := &TreeNode{
		// Type: typ
		// FileId:   id,
		// Name: name,
		//ElementCount: 0,
		Items: make([]*TreeNode, 0),
	}
	return o
}

//--------------------------------

type Tree struct {
	Root  string      `json:"root"`
	Items []*TreeNode `json:"items"`
}
