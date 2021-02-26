// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//CompareStatus ***
type CompareStatus struct {
	Id         int64  `json:"compareId"`
	Priority   int64  `json:"priority"`
	Status     string `json:"status"`
	Thumbnail  string `json:"thumbnail"` // thumbnail http links
	Reason     string `json:"reason"`
	CreateTime string `json:"createTime"`
}

//CompareElement ***
type CompareElement struct {
	DiffType           string `json:"diffType"` // NEW|DELETE|CHANGE,
	Name               string `json:"name"`     // 柱1
	PreviousFileId     int64  `json:"previousFileId"`
	PreviousElementId  int64  `json:"previousElementId"`
	FollowingFileId    int64  `json:"followingFileId"`
	FollowingElementId int64  `json:"followingElementId"`
}

//CompareCategory ***
type CompareCategory struct {
	Id        string `json:"categoryId"`
	Name      string `json:"categoryName"`
	ItemCount int64  `json:"itemCount"`
	Elements  []*CompareElement
}

//CompareData ***
type CompareData struct {
	SpecialtyId   string `json:"specialtyId"`
	SpecialtyName string `json:"specialtyName"`
	ItemCount     int64  `json:"itemCount"`
	Categories    []*CompareCategory
}
