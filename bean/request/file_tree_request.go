// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

type FileTreeRequest struct {
	DesiredHierarchy   []string          `json:"desiredHierarchy,omitempty"`
	CustomizedNodeKeys map[string]string `json:"customizedNodeKeys,omitempty"`
}

func NewFileTreeRequest() *FileTreeRequest {
	o := &FileTreeRequest{
		DesiredHierarchy:   nil,
		CustomizedNodeKeys: nil,
	}
	return o
}
