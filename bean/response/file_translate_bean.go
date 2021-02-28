// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type FileTranslateBean struct {
	FileId     int64    `json:"fileId"`
	Name       string   `json:"name"`
	DatabagId  string   `json:"databagId"`
	Status     string   `json:"status"`
	Reason     string   `json:"reason"`
	Thumbnail  []string `json:"thumbnail"` // thumbnail http links
	CreateTime string   `json:"createTime"`
}
