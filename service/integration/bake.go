// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE integrate.

package integration

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const (
	//为集成模型创建bake数据包 PUT https://api.bimface.com/integrations/{integrateId}/bake
	bakeURI = "integrations/%d/bake"
)

func (o *Service) bakeURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+bakeURI, integrateId)
}

//------------------------------------------------------------------

func (o *Service) Bake(integrateId int64, config map[string]string) (*response.DataBagBean, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := make(map[string]interface{})
	body["config"] = config

	resp := o.ServiceClient.Put(o.bakeURL(integrateId), headers.Header, req.BodyJSON(body))

	var result *response.DataBagBean
	err = utils.RespToBean(resp, result)

	return result, err
}

func (o *Service) DefaultBake(integrateId int64) (*response.DataBagBean, error) {

	body := make(map[string]string)
	body["texture"] = "true"

	return o.Bake(integrateId, body)
}

//---------------------------------------

//查询bake数据包
func (o *Service) GetBakeDataBag(integrateId int64) ([]*response.DataBagBean, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.bakeURL(integrateId), headers.Header)

	var result []*response.DataBagBean
	err = utils.RespToBean(resp, &result)

	return result, err
}
