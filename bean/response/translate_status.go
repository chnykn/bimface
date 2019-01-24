// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//TranslateStatus ***
type TranslateStatus struct {
	FileID     int64
	Name       string
	Priority   int64
	Status     string
	Thumbnail  []string // thumbnail http links
	Reason     string
	CreateTime string
}

//NewTranslateStatus ***
func NewTranslateStatus() *TranslateStatus { //fileID int64, name string
	o := &TranslateStatus{
		// FileID:   fileID,
		// Name: name,
		Thumbnail: make([]string, 0),
	}
	return o
}
