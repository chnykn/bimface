// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type LinkGraphNode struct {
	Name      string `json:"name"`
	FileId    string `json:"fileId"`
	DatabagId string `json:"databagId"`

	LinkName      string `json:"linkName"`
	LinkPathHash  string `json:"linkPathHash"`
	LinkTransform string `json:"linkTransform"`

	Links  []*LinkGraphNode `json:"links,omitempty"`
	Params []interface{}    `json:"params,omitempty"`
}

type LinkGraphNodeBean LinkGraphNode
