// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package request

type IntegrateQueryRequest struct {
	AppKey string `json:"appKey,omitempty"` //应用的appKey 可选

	FileName      string `json:"fileName,omitempty"` //单模型的名称 可选
	SourceId      string `json:"sourceId,omitempty"`
	IntegrateId   string `json:"integrateId,omitempty"`
	IntegrateType string `json:"integrateType,omitempty"`

	PageNo   int `json:"pageNo,omitempty"`   //页码 可选
	PageSize int `json:"pageSize,omitempty"` //应用的appKey 可选

	Status   int    `json:"status,omitempty"`   //模型状态码：1(处理中); 99(成功); -1(失败； 可选
	SortType string `json:"sortType,omitempty"` //筛选类型 示例：create_time desc 可选

	StartDate string `json:"startDate,omitempty"` //开始日期 示例：2019-05-01 desc 可选
	EndDate   string `json:"endDate,omitempty"`   //截止日期 示例：2019-05-01 desc 可选
}

func NewIntegrateQueryRequest() *IntegrateQueryRequest {
	o := &IntegrateQueryRequest{}
	return o
}
