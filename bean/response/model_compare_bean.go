// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/v3/bean/common"

type ModelCompareBean struct {
	CompareId int64 `json:"compareId"`

	Name     string `json:"name"`     //集成模型的名称
	SourceId string `json:"sourceId"` //模型对应的sourceld

	OfflineDatabagStatus string `json:"offlineDatabagStatus"`

	Reason string `json:"reason"` //若转换失败，返回失败原因
	Status string `json:"status"`

	Type       string   `json:"type"`       //
	WorkerType string   `json:"workerType"` //
	Thumbnail  []string `json:"thumbnail"`  //模型的缩略图

	Cost       int    `json:"cost"`       //任务耗时(秒)
	CreateTime string `json:"createTime"` //创建时间 示例：2019-05-01 desc
}

type ModelCompareBeanPageList struct {
	common.PageList
	List []*ModelCompareBean `json:"list"`
}
