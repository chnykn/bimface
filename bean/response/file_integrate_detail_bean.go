// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/v2/bean/common"

type FileIntegrateDetailBean struct {
	IntegrateId int64  `json:"integrateId"`
	SourceId    string `json:"sourceId"` //模型对应的sourceld
	Name        string `json:"name"`     //集成模型的名称

	DatabagId            string `json:"databagId"` //数据包id
	OfflineDatabagStatus string `json:"offlineDatabagStatus"`

	ShareToken string `json:"shareToken"`
	ShareUrl   string `json:"shareUrl"`

	Reason string `json:"reason"` //若转换失败，返回失败原因
	Status string `json:"status"`

	Type       string   `json:"type"`       //转换类型 rvt-translate(或者igms-translate…​)
	WorkerType string   `json:"workerType"` //转换类型 rvt-translate(或者igms-translate…​)
	Thumbnail  []string `json:"thumbnail"`  //模型的缩略图

	Cost       int    `json:"cost"`       //任务耗时(秒)
	CreateTime string `json:"createTime"` //创建时间 示例：2019-05-01 desc
}

type FileIntegrateDetailBeanPageList struct {
	common.PageList
	List []*FileIntegrateDetailBean `json:"list"`
}
