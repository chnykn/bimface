// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/bean/common"

type FileTranslateDetailBean struct {
	AppKey               string   `json:"appKey"`               //应用的appKey
	Cost                 int      `json:"cost"`                 //任务耗时(秒)
	CreateTime           string   `json:"createTime"`           //创建时间 示例：2019-05-01 desc
	DatabagId            string   `json:"databagId"`            //数据包id
	FileId               int64    `json:"fileId"`               //单模型对应的id
	Length               int64    `json:"length"`               //文件长度
	Name                 string   `json:"name"`                 //集成模型的名称
	OfflineDatabagStatus string   `json:"offlineDatabagStatus"` //集成模型的名称
	Reason               string   `json:"reason"`               //若转换失败，返回失败原因
	Retry                bool     `json:"retry"`                //重试
	ShareToken           string   `json:"shareToken"`           //分享码
	ShareUrl             string   `json:"shareUrl"`             //分享链接
	SourceId             string   `json:"sourceId"`             //模型对应的sourceld
	Status               string   `json:"status"`               //模型状态 processing(处理中); success(成功); failed(失败)
	SupprtOfflineDatabag bool     `json:"supprtOfflineDatabag"` //是否支持离线数据包
	Thumbnail            []string `json:"thumbnail"`            //模型的缩略图
	Type                 string   `json:"type"`                 //转换类型 rvt-translate(或者igms-translate…​)
}

type FileTranslateDetailBeanPageList struct {
	common.PageList
	List []*FileTranslateDetailBean `json:"list"`
}
