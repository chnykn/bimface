// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
	"github.com/imroc/req"
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

func (o *PropertyService) elemPropertyURL(fileID int64, elementID string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+propertyURI, fileID, elementID)
}

func (o *PropertyService) elemPropertiesURL(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+propertiesURI, fileID)
}

func (o *PropertyService) integratePropertyURL(integrateID int64, elementID string) string { //fileID int64,
	result := fmt.Sprintf(o.Endpoint.APIHost+integratePropertyURI2, integrateID, elementID)
	//if includeOverrides {
	//	result = result + "?includeOverrides=true"
	//}
	return result
}

//-----------------------------------------------------------------------------------

//GetElementPropertyResp ***
func (o *PropertyService) GetElementPropertyResp(fileID int64, elementID string) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.elemPropertyURL(fileID, elementID), headers.Header)
	return resp, err
}

//GetElementProperty 文件转换相关: 获取文件转换的构件属性
//http://static.bimface.com/restful-apidoc/dist/translateSingleModel.html#_getelementusingget_1
/***
字段		类型	必填	描述
fileId		Number	Y	文件ID
elementId	String	Y	构件ID
***/
func (o *PropertyService) GetElementProperty(fileID int64, elementID string) (*response.Element, error) {
	resp, err := o.GetElementPropertyResp(fileID, elementID)
	if err != nil {
		return nil, err
	}

	result := response.NewElement(elementID)
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//TODO ....

//GetElementPropertyResp ***
func (o *PropertyService) GetElementPropertiesResp(fileID int64) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Post(o.elemPropertiesURL(fileID), headers.Header, `{ "elementIds" : [ "674160", "688045" ]}`)
	return resp, err
}

func (o *PropertyService) GetElementProperties(fileID int64) ([]*response.Element, error) {
	resp, err := o.GetElementPropertiesResp(fileID)
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
func (o *PropertyService) GetIntgrElementPropertyResp(integrateID, fileID int64, elementID string) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integratePropertyURL(integrateID, fileID, elementID), headers.Header)
	return resp, err
}
*/

//GetIntgrElementProperty 模型集成相关: 获取集成模型的构件属性
//http://static.bimface.com/book/restful/articles/api/integrate/get-integrate-element-prop.html
/***
字段		类型	必填	描述
fileId		Number	Y	文件ID
elementId	String	Y	构件ID
***/
/*
func (o *PropertyService) GetIntgrElementProperty(integrateID, fileID int64, elementID string) (*response.Element, error) {
	resp, err := o.GetIntgrElementPropertyResp(integrateID, fileID, elementID)
	if err != nil {
		return nil, err
	}

	result := response.NewElement(elementID)
	err = utils.RespToBean(resp, result)

	return result, err
}
*/

//-----------------------------------------------------------------------------------

//GetIntgrElementPropertyResp ***
func (o *PropertyService) GetIntgrElementPropertyResp(integrateID int64, elementID string) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integratePropertyURL(integrateID, elementID), headers.Header)
	return resp, err
}

//GetIntgrElementProperty 模型集成相关: 获取集成模型的构件属性
//http://static.bimface.com/book/restful/articles/api/integrate/get-integrate-element-prop.html
/***
字段		类型	必填	描述
fileId		Number	Y	文件ID
elementId	String	Y	构件ID
***/
func (o *PropertyService) GetIntgrElementProperty(integrateID int64, elementID string) (*response.Element, error) {
	resp, err := o.GetIntgrElementPropertyResp(integrateID, elementID)
	if err != nil {
		return nil, err
	}

	result := response.NewElement(elementID)
	err = utils.RespToBean(resp, result)

	return result, err
}
