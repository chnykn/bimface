// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//
type FileBean struct {
	FileId     int64  `json:"fileId"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Etag       string `json:"etag"`
	Suffix     string `json:"suffix"`
	Length     int64  `json:"length"`
	CreateTime string `json:"createTime"`
}

func (o *FileBean) ToString() string {
	return fmt.Sprintf("FileBean [FileId=%d, Name=%s, Status=%s, ETag=%s, Suffix=%s, Length=%d, CreateTime=%s]",
		o.FileId, o.Name, o.Status, o.Etag, o.Suffix, o.Length, o.CreateTime)
}
