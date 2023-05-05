// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

type Page struct {
	StartIndex int `json:"startIndex"` //起始索引数
	PageNo     int `json:"pageNo"`     //当前页码
	PrePage    int `json:"prePage "`   //上一页码
	NextPage   int `json:"nextPage"`   //下一页码

	PageSize   int `json:"pageSize"`    //每页条目数
	TotalCount int `json:"totalCount "` //条目总数
	TotalPages int `json:"totalPages "` //页码总数

	HtmlDisplay int `json:"htmlDisplay "` //
}

type PageList struct {
	Page *Page `json:"page"`
	//List []interface{} `json:"list"` 等待泛型
}
