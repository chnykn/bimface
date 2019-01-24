// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//FileBean ***
type FileBean struct {
	FileID     int64
	Name       string
	Status     string
	ETag       string
	Suffix     string
	Length     int64
	CreateTime string
}

//NewFileBean ***
func NewFileBean() *FileBean { //id int64, name string
	o := &FileBean{
		// ID:   id,
		// Name: name,
	}
	return o
}

// ToString get the string
func (o *FileBean) ToString() string {
	return fmt.Sprintf("FileBean [FileID=%d, Name=%s, Status=%s, ETag=%s, Suffix=%s, Length=%d, CreateTime=%s]",
		o.FileID, o.Name, o.Status, o.ETag, o.Suffix, o.Length, o.CreateTime)
}
