// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/http"
	"github.com/chnykn/bimface/utils"
	"fmt"
)

const (
	getElevsURI            string = "/data/v2/files/%d/floors"
	getIntegrationElevsURI string = "/data/v2/integrations/%d/floors"
)

//ElevService ***
type ElevService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewElevService ***
func NewElevService(serviceClient *http.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *ElevService {
	o := &ElevService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //http.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *ElevService) getElevsURL(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getElevsURI, fileID)
}

func (o *ElevService) getIntegrationElevsURL(integrationID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getIntegrationElevsURI, integrationID)
}

//---------------------------------------------------------------------

//GetElevs 获取文件转换的楼层信息
//http://static.bimface.com/book/restful/articles/api/translate/get-floors.html
func (o *ElevService) GetElevs(fileID int64) ([]response.Elev, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getElevsURL(fileID), headers.Header)

	result, err := http.RespToBeans(resp, &response.Elev{})
	return result.([]response.Elev), nil
}

//GetIntegrationElevs 获取集成模型楼层信息
//http://static.bimface.com/book/restful/articles/api/integrate/get-integrate-floors.html
func (o *ElevService) GetIntegrationElevs(integrateID int64) ([]response.Elev, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getIntegrationElevsURL(integrateID), headers.Header)

	result, err := http.RespToBeans(resp, &response.Elev{})
	return result.([]response.Elev), nil
}
