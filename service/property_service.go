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
	//获取文件转换的构件属性
	propertyURI string = "/data/element/property?fileId=%d&elementId=%s"

	integratePropertyURI string = "/data/integration/element/property?integrateId=%d&fileId=%d&elementId=%s"
)

//PropertyService ***
type PropertyService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewPropertyService ***
func NewPropertyService(serviceClient *http.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *PropertyService {
	o := &PropertyService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //http.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *PropertyService) propertyURL(fileID int64, elementID string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+propertyURI, fileID, elementID)
}

func (o *PropertyService) integratePropertyURL(integrateID, fileID int64, elementID string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+integratePropertyURI, integrateID, fileID, elementID)
}

//---------------------------------------------------------------------

//GetElementProperty 文件转换相关: 获取文件转换的构件属性
//http://static.bimface.com/book/restful/articles/api/translate/get-ele-prop.html
/***
字段		类型	必填	描述
fileId		Number	Y	文件ID
elementId	String	Y	构件ID
***/
func (o *PropertyService) GetElementProperty(fileID int64, elementID string) (*response.PropertyPack, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.propertyURL(fileID, elementID), headers.Header)

	result := response.NewPropertyPack(elementID)
	err = http.RespToBean(resp, result)

	return result, err
}

//GetIntegrateElementProperty 模型集成相关: 获取集成模型的构件属性
//http://static.bimface.com/book/restful/articles/api/integrate/get-integrate-element-prop.html
/***
字段		类型	必填	描述
fileId		Number	Y	文件ID
elementId	String	Y	构件ID
***/
func (o *PropertyService) GetIntegrateElementProperty(integrateID, fileID int64, elementID string) (*response.PropertyPack, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integratePropertyURL(integrateID, fileID, elementID), headers.Header)

	result := response.NewPropertyPack(elementID)
	err = http.RespToBean(resp, result)

	return result, err
}
