// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package request

//TranslateDetailRequest ***
type TranslateDetailRequest struct {
	AppKey    string `json:"appKey,omitempty"`    //应用的appKey 可选
	FileId    string `json:"fileId,omitempty"`    //单模型对应的id 可选
	Suffix    string `json:"suffix,omitempty"`    //单模型的文件类型	rvt(或者igms,dwg…) 可选
	FileName  string `json:"fileName,omitempty"`  //单模型的名称 可选
	SourceId  string `json:"sourceId,omitempty"`  //模型对应的sourceId 可选
	PageNo    *int   `json:"pageNo,omitempty"`    //页码 可选
	PageSize  *int   `json:"pageSize,omitempty"`  //应用的appKey 可选
	Status    *int   `json:"status,omitempty"`    //模型状态码：1(处理中); 99(成功); -1(失败； 可选
	SortType  string `json:"sortType,omitempty"`  //筛选类型 示例：create_time desc 可选
	StartDate string `json:"startDate,omitempty"` //开始日期 示例：2019-05-01 desc 可选
	EndDate   string `json:"endDate,omitempty"`   //截止日期 示例：2019-05-01 desc 可选
}

//NewTranslateDetailRequest ***
func NewTranslateDetailRequest(sourceFileId int64) *TranslateDetailRequest {
	o := &TranslateDetailRequest{
		AppKey:    "",
		FileId:    "",
		Suffix:    "",
		FileName:  "",
		SourceId:  "",
		PageNo:    nil,
		PageSize:  nil,
		Status:    nil,
		SortType:  "",
		StartDate: "",
		EndDate:   "",
	}
	return o
}
