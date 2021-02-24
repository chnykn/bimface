// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

const (
	getFloorsURI            string = "/data/v2/files/%d/floors"
	getIntegrationFloorsURI string = "/data/v2/integrations/%d/floors"
)

//FloorService ***
type FloorService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewFloorService ***
func NewFloorService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *FloorService {
	o := &FloorService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *FloorService) getFloorsURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getFloorsURI, fileId)
}

func (o *FloorService) getIntegrationFloorsURL(integrationId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getIntegrationFloorsURI, integrationId)
}

//-----------------------------------------------------------------------------------

//GetFloorsResp ***
func (o *FloorService) GetFloorsResp(fileId int64) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getFloorsURL(fileId), headers.Header)
	return resp, nil
}

//GetFloors 获取文件转换的楼层信息
//http://static.bimface.com/book/restful/articles/api/translate/get-floors.html
func (o *FloorService) GetFloors(fileId int64) ([]*response.Floor, error) {
	resp, err := o.GetFloorsResp(fileId)
	if err != nil {
		return nil, err
	}

	result := make([]*response.Floor, 0)
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//-----------------------------------------------------------------------------------

//GetIntegrationFloorsResp ***
func (o *FloorService) GetIntegrationFloorsResp(integrateId int64) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getIntegrationFloorsURL(integrateId), headers.Header)
	return resp, nil
}

//GetIntegrationFloors 获取集成模型楼层信息
//http://static.bimface.com/book/restful/articles/api/integrate/get-integrate-floors.html
func (o *FloorService) GetIntegrationFloors(integrateId int64) ([]*response.Floor, error) {
	resp, err := o.GetIntegrationFloorsResp(integrateId)
	if err != nil {
		return nil, err
	}

	result := make([]*response.Floor, 0)
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
