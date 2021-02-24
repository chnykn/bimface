// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//TranslateStatus ***
type TranslateStatus struct {
	FileId     int64    `json:"fileId"`
	Name       string   `json:"name"`
	Status     string   `json:"status"`
	Reason     string   `json:"reason"`
	Thumbnail  []string `json:"thumbnail"` // thumbnail http links
	CreateTime string   `json:"createTime"`
}

//NewTranslateStatus ***
func NewTranslateStatus() *TranslateStatus { //fileId int64, name string
	o := &TranslateStatus{
		// FileId:   fileId,
		// Name: name,
		Thumbnail: make([]string, 0),
	}
	return o
}
