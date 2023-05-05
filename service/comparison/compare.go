// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comparison

import (
	"strconv"

	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/httpkit"
)

const (
	//发起模型对比 POST https://api.bimface.com/v2/compare
	compareURI string = "v2/compare"
)

// 发起模型对比
func (o *Service) Compare(compareRequest *request.CompareRequest) (*response.ModelCompareBean, error) {

	result := new(response.ModelCompareBean)

	url := o.Endpoint.APIHost + compareURI
	body := httpkit.JsonReqBody(compareRequest)

	err := o.POST(url, result, body)
	return result, err
}

// 获取模型对比状态
func (o *Service) GetStatus(compareId int64) (*response.ModelCompareBean, error) {
	result := new(response.ModelCompareBean)

	url := o.Endpoint.APIHost + compareURI + "?compareId=" + strconv.FormatInt(compareId, 10)
	err := o.GET(url, result)

	return result, err
}
