// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//CompareStatus ***
type CompareStatus struct {
	ID         int64  `json:"compareId"`
	Priority   int64  `json:"priority"`
	Status     string `json:"status"`
	Thumbnail  string `json:"thumbnail"` // thumbnail http links
	Reason     string `json:"reason"`
	CreateTime string `json:"createTime"`
}

//NewCompareStatus ***
func NewCompareStatus() *CompareStatus { //fileID int64, name string
	o := &CompareStatus{
		// FileID:   fileID,
		// Name: name,
		//Thumbnail: make([]string, 0),
	}
	return o
}

//=================================================================

//CompareElement ***
type CompareElement struct {
	DiffType           string `json:"diffType"` // NEW|DELETE|CHANGE,
	Name               string `json:"name"`     // 柱1
	PreviousFileID     int64  `json:"previousFileId"`
	PreviousElementID  int64  `json:"previousElementId"`
	FollowingFileID    int64  `json:"followingFileId"`
	FollowingElementID int64  `json:"followingElementId"`
}

//----------------------------------------------------------

//CompareCategory ***
type CompareCategory struct {
	ID        string `json:"categoryId"`
	Name      string `json:"categoryName"`
	ItemCount int64  `json:"itemCount"`
	Elements  []*CompareElement
}

//NewCompareCategory ***
func NewCompareCategory() *CompareCategory {
	o := &CompareCategory{
		// CategoryID:   ,
		// CategoryName: ,
		Elements: make([]*CompareElement, 0),
	}
	return o
}

//----------------------------------------------------------

//CompareData ***
type CompareData struct {
	SpecialtyID   string `json:"specialtyId"`
	SpecialtyName string `json:"specialtyName"`
	ItemCount     int64  `json:"itemCount"`
	Categories    []*CompareCategory
}

//NewCompareData ***
func NewCompareData() *CompareData {
	o := &CompareData{
		// SpecialtyID:   ,
		// SpecialtyName: ,
		Categories: make([]*CompareCategory, 0),
	}
	return o
}
