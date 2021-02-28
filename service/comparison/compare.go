// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comparison

import (
	"strconv"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/v2/bean/request"
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	//发起模型对比 POST https://api.bimface.com/v2/compare
	compareURI string = "v2/compare"
)

//发起模型对比
func (o *Service) Compare(compareRequest *request.CompareRequest) (*response.ModelCompareBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	url := o.Endpoint.APIHost + compareURI
	body := req.BodyJSON(compareRequest)
	resp := o.ServiceClient.Post(url, headers.Header, body)

	var result *response.ModelCompareBean
	err = utils.RespToBean(resp, result)

	return result, err
}

//获取模型对比状态
func (o *Service) GetStatus(compareId int64) (*response.ModelCompareBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	url := o.Endpoint.APIHost + compareURI + "?compareId=" + strconv.FormatInt(compareId, 10)
	resp := o.ServiceClient.Get(url, headers.Header)

	var result *response.ModelCompareBean
	err = utils.RespToBean(resp, result)

	return result, err
}
