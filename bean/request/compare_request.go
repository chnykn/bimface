// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

type CompareRequest struct {
	Name     string `json:"Name"`               //用户指定对比后的模型的名字
	SourceId string `json:"sourceId,omitempty"` //第三方应用自己的ID

	ComparedEntityType string `json:"comparedEntityType"` //"file"或"integration"
	PreviousId         int64  `json:"previousId"`         //变更前文件ID
	FollowingId        int64  `json:"followingId"`        //变更后文件ID

	Config   map[string]string `json:"config"` //发起模型对比支持的配置
	Callback string            `json:"callback"`
}

func NewCompareRequest(entityType string, previousId, followingId int64) *CompareRequest {
	o := &CompareRequest{
		ComparedEntityType: entityType,
		PreviousId:         previousId,
		FollowingId:        followingId,
		Config:             nil,
	}
	return o
}
