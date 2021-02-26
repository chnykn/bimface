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

	//integratePropertyURI string = "/data/integration/element/property?integrateId=%d&fileId=%d&elementId=%s"

	//GET https://api.bimface.com/data/v2/integrations/{integrateId}/elements/{elementId}
	integratePropertyURI2 string = "/data/v2/integrations/%d/elements/%s"
)

//PropertyService ***
type PropertyService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewPropertyService ***
func NewPropertyService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *PropertyService {
	o := &PropertyService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *PropertyService) integratePropertyURL(integrateId int64, elementId string) string { //fileId int64,
	result := fmt.Sprintf(o.Endpoint.APIHost+integratePropertyURI2, integrateId, elementId)
	//if includeOverrides {
	//	result = result + "?includeOverrides=true"
	//}
	return result
}

//-----------------------------------------------------------------------------------

//GetIntgrElementPropertyResp ***
func (o *PropertyService) GetIntgrElementPropertyResp(integrateId int64, elementId string) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integratePropertyURL(integrateId, elementId), headers.Header)
	return resp, err
}

//GetIntgrElementProperty 模型集成相关: 获取集成模型的构件属性
//http://static.bimface.com/book/restful/articles/api/integrate/get-integrate-element-prop.html
/***
字段		类型	必填	描述
fileId		Number	Y	文件Id
elementId	String	Y	构件Id
***/
func (o *PropertyService) GetIntgrElementProperty(integrateId int64, elementId string) (*response.Element, error) {
	resp, err := o.GetIntgrElementPropertyResp(integrateId, elementId)
	if err != nil {
		return nil, err
	}

	result := response.NewElement(elementId)
	err = utils.RespToBean(resp, result)

	return result, err
}
