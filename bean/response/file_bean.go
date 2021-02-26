// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//FileBean ***
type FileBean struct {
	Id         int64  `json:"fileId"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Etag       string `json:"etag"`
	Suffix     string `json:"suffix"`
	Length     int64  `json:"length"`
	CreateTime string `json:"createTime"`
}

// ToString get the string
func (o *FileBean) ToString() string {
	return fmt.Sprintf("FileBean [Id=%d, Name=%s, Status=%s, ETag=%s, Suffix=%s, Length=%d, CreateTime=%s]",
		o.Id, o.Name, o.Status, o.Etag, o.Suffix, o.Length, o.CreateTime)
}
