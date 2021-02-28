// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/bean"
	"github.com/chnykn/bimface/bean/request"
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const (
	//发起模型集成 PUT    https://api.bimface.com/integrate
	integrateURI string = "/integrate"

	//查询集成状态 GET    https://api.bimface.com/integrate
	//删除集成模型 DELETE https://api.bimface.com/integrate
	integrateWithIdURI string = "/integrate?integrateId=%d"

	//批量查询集成模型状态详情 POST https://api.bimface.com/integrateDetails
	integrateDetailsURI string = "/integrateDetails"
)

//---------------------------------------------------------------------

func (o *Service) integrateURL() string {
	return o.Endpoint.APIHost + integrateURI
}

func (o *Service) integrateWithIdURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+integrateWithIdURI, integrateId)
}

func (o *Service) integrateDetailsURL() string {
	return fmt.Sprintf(o.Endpoint.APIHost + integrateDetailsURI)
}

//---------------------------------------------------------------------

func (o *Service) Integrate(integrateRequest *request.FileIntegrateRequest) (*response.FileIntegrateBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := req.BodyJSON(integrateRequest)
	resp := o.ServiceClient.Put(o.integrateURL(), headers.Header, body)

	var result *response.FileIntegrateBean
	err = utils.RespToBean(resp, result)

	return result, err
}

//获取转换状态
func (o *Service) GetStatus(integrateId int64) (*response.FileIntegrateBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integrateWithIdURL(integrateId), headers.Header)

	var result *response.FileIntegrateBean
	err = utils.RespToBean(resp, result)

	return result, err
}

//获取转换状态
func (o *Service) Delete(integrateId int64) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Delete(o.integrateWithIdURL(integrateId), headers.Header)

	var result *bean.RespResult
	result, err = utils.RespToResult(resp)

	var ret string
	if result != nil {
		ret = result.Code
	}

	return ret, nil
}

//---------------------------------------------------------------------

//批量查询集成模型状态详情
func (o *Service) GetDetails(queryRequest *request.IntegrateQueryRequest) (*response.FileIntegrateDetailBeanPageList, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	var resp *req.Resp
	if queryRequest != nil {
		body := req.BodyJSON(queryRequest)
		resp = o.ServiceClient.Post(o.integrateDetailsURL(), body, headers.Header)
	} else {
		resp = o.ServiceClient.Post(o.integrateDetailsURL(), headers.Header)
	}

	var result *response.FileIntegrateDetailBeanPageList
	err = utils.RespToBean(resp, result)

	return result, err
}
