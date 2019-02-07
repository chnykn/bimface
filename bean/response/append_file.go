// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//AppendFile ***
type AppendFile struct {
	ID         int64  `json:"appendFileId"`
	Name       string `json:"name"`
	Length     int64  `json:"length"`
	Position   int64  `json:"position"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
	File       FileBean
}

//NewAppendFile ***
func NewAppendFile() *AppendFile {
	o := &AppendFile{
		//File: NewFileBean(),
	}
	return o
}
