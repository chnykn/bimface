// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type SplitDrawingStatus struct {
	FileId         int64  `json:"fileId"` //单模型对应的id
	Length         int64  `json:"length"` //文件长度
	Status         string `json:"status"`
	Reason         string `json:"reason"`
	DatabagVersion string `json:"databagVersion"` //数据包id
	CreateTime     string `json:"createTime"`
}

//NewSplitDrawingStatus ***
func NewSplitDrawingStatus() *SplitDrawingStatus { //fileId int64,
	o := &SplitDrawingStatus{
		// FileId   fileId,
		// Name: name,
	}
	return o
}
