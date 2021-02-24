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
	//获取文件转换的构件列表
	elementIdsURI string = "/data/v2/files/%d/elementIds" //data/v2/files/{fileId}/elementIds

	//获取集成模型的构件列表
	intgrElementURI string = "/data/integration/element?integrateId=%d"
)

//ElementService ***
type ElementService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewElementService ***
func NewElementService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *ElementService {
	o := &ElementService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *ElementService) elementIdsURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+elementIdsURI, fileId)
}

func (o *ElementService) intgrElementURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+intgrElementURI, integrateId)
}

//-----------------------------------------------------------------------------------

//GetElementsResp *** http://static.bimface.com/restful-apidoc/dist/translateSingleModel.html#_getelementidsusingget
func (o *ElementService) GetElementsResp(fileId int64, params req.QueryParam) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.elementIdsURL(fileId), params, headers.Header)
	return resp, nil
}

//GetElementsWithParams 文件转换相关: 获取文件转换的构件列表
//必填参数: fileId  params相关参数，详见 http://static.bimface.com/restful-apidoc/dist/translateSingleModel.html#_getelementidsusingget
func (o *ElementService) GetElementsWithParams(fileId int64, params req.QueryParam) ([]string, error) {
	resp, err := o.GetElementsResp(fileId, params)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//GetElements 文件转换相关: 获取文件转换的构件列表
//http://static.bimface.com/restful-apidoc/dist/translateSingleModel.html#_getelementidsusingget
/***
字段		类型	必填	描述
fileId		Number	Y	文件Id
specialty	String	N	专业
floor		String	N	楼层
categoryId	String	N	构件分类Id
family		String	N	族
familyType	String	N	族类型
*** 其他参数详见网址地址
***/
func (o *ElementService) GetElements(fileId int64, floor, specialty, categoryId,
	family, familyType string) ([]string, error) {

	params := make(req.QueryParam)
	if floor != "" {
		params["floor"] = floor
	}
	if specialty != "" {
		params["specialty"] = specialty
	}
	if categoryId != "" {
		params["categoryId"] = categoryId
	}
	if family != "" {
		params["family"] = family
	}
	if familyType != "" {
		params["familyType"] = familyType
	}

	return o.GetElementsWithParams(fileId, params)
}

//-----------------------------------------------------------------------------------

//GetIntgrElementsResp ***
func (o *ElementService) GetIntgrElementsResp(integrateId int64, params req.QueryParam) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.intgrElementURL(integrateId), params, headers.Header)
	return resp, nil
}

//GetIntgrElementsWithParams ***
func (o *ElementService) GetIntgrElementsWithParams(integrateId int64, params req.QueryParam) (*response.IntgrElements, error) {
	resp, err := o.GetIntgrElementsResp(integrateId, params)
	if err != nil {
		return nil, err
	}

	result := response.NewElements()
	err = utils.RespToBean(resp, result)

	return result, err
}

//GetIntgrElements 模型集成相关: 获取集成的构件列表
//http://doc.bimface.com/book/restful/articles/api/integrate/get-integrate-element.html
/***
字段		类型	必填	描述
integrateId	Number	Y	集成Id
specialty	String	N	专业
floor		String	N	楼层
categoryId	String	N	构件分类Id
family		String	N	族
familyType	String	N	族类型
***/
func (o *ElementService) GetIntgrElements(fileId int64, floor, specialty, categoryId,
	family, familyType string) (*response.IntgrElements, error) {

	params := make(req.QueryParam)
	if floor != "" {
		params["floor"] = floor
	}
	if specialty != "" {
		params["specialty"] = specialty
	}
	if categoryId != "" {
		params["categoryId"] = categoryId
	}
	if family != "" {
		params["family"] = family
	}
	if familyType != "" {
		params["familyType"] = familyType
	}

	return o.GetIntgrElementsWithParams(fileId, params)
}
