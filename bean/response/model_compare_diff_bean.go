// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type ModelCompareDiff struct {
	Id string `json:"Id"`

	DiffType    string `json:"diffType"`    //只有三种：NEW DELETE CHANGE
	PreviousId  int64  `json:"previousId"`  //变更前文件ID
	FollowingId int64  `json:"followingId"` //变更后文件ID

	//------------

	Specialty string `json:"specialty"`

	CategoryId   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`

	Family string `json:"family"`

	ElementId   string `json:"elementId"`
	ElementName string `json:"elementName"`
}

type ModelCompareDiffBean ModelCompareDiff

//----------------------------------------

type ModelCompareDiffPagination struct {
	Page  int                     `json:"page"`
	Total int                     `json:"total"`
	Data  []*ModelCompareDiffBean `json:"data"`
}

type ModelCompareDiffsBean ModelCompareDiffPagination
