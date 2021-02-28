// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type DrawingCompareDiff struct {
	Id       string `json:"Id"`
	DiffType string `json:"diffType"` //只有三种：NEW DELETE CHANGE
	Layer    string `json:"layer"`
}

type DrawingCompareDiffBean DrawingCompareDiff

//----------------------------------------

type DrawingCompareDiffPagination struct {
	Page  int                       `json:"page"`
	Total int                       `json:"total"`
	Data  []*DrawingCompareDiffBean `json:"data"`
}

type DrawingCompareDiffsBean DrawingCompareDiffPagination
