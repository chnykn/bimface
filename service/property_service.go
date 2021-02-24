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
	//获取文件转换的构件属性
	propertyURI string = "/data/v2/files/%d/elements/%s" //data/v2/files/{fileId}/elements/{elementId}

	//批量获取构件属性 POST https://api.bimface.com/data/v2/files/{fileId}/elements
	propertiesURI string = "/data/v2/files/%d/elements" //data/v2/files/{fileId}/elements

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

func (o *PropertyService) elemPropertyURL(fileId int64, elementId string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+propertyURI, fileId, elementId)
}

func (o *PropertyService) elemPropertiesURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+propertiesURI, fileId)
}

func (o *PropertyService) integratePropertyURL(integrateId int64, elementId string) string { //fileId int64,
	result := fmt.Sprintf(o.Endpoint.APIHost+integratePropertyURI2, integrateId, elementId)
	//if includeOverrides {
	//	result = result + "?includeOverrides=true"
	//}
	return result
}

//-----------------------------------------------------------------------------------

//GetElementPropertyResp ***
func (o *PropertyService) GetElementPropertyResp(fileId int64, elementId string) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.elemPropertyURL(fileId, elementId), headers.Header)
	return resp, err
}

//GetElementProperty 文件转换相关: 获取文件转换的构件属性
//http://static.bimface.com/restful-apidoc/dist/translateSingleModel.html#_getelementusingget_1
/***
字段		类型	必填	描述
fileId		Number	Y	文件Id
elementId	String	Y	构件Id
***/
func (o *PropertyService) GetElementProperty(fileId int64, elementId string) (*response.Element, error) {
	resp, err := o.GetElementPropertyResp(fileId, elementId)
	if err != nil {
		return nil, err
	}

	result := response.NewElement(elementId)
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//TODO ....

//GetElementPropertyResp ***
func (o *PropertyService) GetElementPropertiesResp(fileId int64) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Post(o.elemPropertiesURL(fileId), headers.Header, `{ "elementIds" : [ "674160", "688045" ]}`)
	return resp, err
}

func (o *PropertyService) GetElementProperties(fileId int64) ([]*response.Element, error) {
	resp, err := o.GetElementPropertiesResp(fileId)
	if err != nil {
		return nil, err
	}

	result := make([]*response.Element, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}

//-----------------------------------------------------------------------------------
/*
//GetIntgrElementPropertyResp ***
func (o *PropertyService) GetIntgrElementPropertyResp(integrateId, fileId int64, elementId string) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integratePropertyURL(integrateId, fileId, elementId), headers.Header)
	return resp, err
}
*/

//GetIntgrElementProperty 模型集成相关: 获取集成模型的构件属性
//http://static.bimface.com/book/restful/articles/api/integrate/get-integrate-element-prop.html
/***
字段		类型	必填	描述
fileId		Number	Y	文件Id
elementId	String	Y	构件Id
***/
/*
func (o *PropertyService) GetIntgrElementProperty(integrateId, fileId int64, elementId string) (*response.Element, error) {
	resp, err := o.GetIntgrElementPropertyResp(integrateId, fileId, elementId)
	if err != nil {
		return nil, err
	}

	result := response.NewElement(elementId)
	err = utils.RespToBean(resp, result)

	return result, err
}
*/

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
