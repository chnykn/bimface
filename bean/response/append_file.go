// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//AppendFile ***
type AppendFile struct {
	AppendFileID int64
	Name         string
	Length       int64
	Position     int64
	Status       string
	CreateTime   string
	File         FileBean
}

//NewAppendFile ***
func NewAppendFile() *AppendFile {
	o := &AppendFile{
		//File: NewFileBean(),
	}
	return o
}
