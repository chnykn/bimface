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
	elementURI string = "/data/element/id?fileId=%d"

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

func (o *ElementService) elementURL(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+elementURI, fileID)
}

func (o *ElementService) intgrElementURL(integrateID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+intgrElementURI, integrateID)
}

//-----------------------------------------------------------------------------------

//GetElementsResp ***
func (o *ElementService) GetElementsResp(fileID int64, params req.QueryParam) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.elementURL(fileID), params, headers.Header)
	return resp, nil
}

//GetElementsWithParams 文件转换相关: 获取文件转换的构件列表
//必填参数: fileID
func (o *ElementService) GetElementsWithParams(fileID int64, params req.QueryParam) ([]string, error) {
	resp, err := o.GetElementsResp(fileID, params)
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
//http://doc.bimface.com/book/restful/articles/api/translate/get-ele-ids.html
/***
字段		类型	必填	描述
fileId		Number	Y	文件ID
specialty	String	N	专业
floor		String	N	楼层
categoryId	String	N	构件分类ID
family		String	N	族
familyType	String	N	族类型
***/
func (o *ElementService) GetElements(fileID int64, floor, specialty, categoryID,
	family, familyType string) ([]string, error) {

	params := make(req.QueryParam)
	if floor != "" {
		params["floor"] = floor
	}
	if specialty != "" {
		params["specialty"] = specialty
	}
	if categoryID != "" {
		params["categoryId"] = categoryID
	}
	if family != "" {
		params["family"] = family
	}
	if familyType != "" {
		params["familyType"] = familyType
	}

	return o.GetElementsWithParams(fileID, params)
}

//-----------------------------------------------------------------------------------

//GetIntgrElementsResp ***
func (o *ElementService) GetIntgrElementsResp(integrateID int64, params req.QueryParam) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.intgrElementURL(integrateID), params, headers.Header)
	return resp, nil
}

//GetIntgrElementsWithParams ***
func (o *ElementService) GetIntgrElementsWithParams(integrateID int64, params req.QueryParam) (*response.IntgrElements, error) {
	resp, err := o.GetIntgrElementsResp(integrateID, params)
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
integrateId	Number	Y	集成ID
specialty	String	N	专业
floor		String	N	楼层
categoryId	String	N	构件分类ID
family		String	N	族
familyType	String	N	族类型
***/
func (o *ElementService) GetIntgrElements(fileID int64, floor, specialty, categoryID,
	family, familyType string) (*response.IntgrElements, error) {

	params := make(req.QueryParam)
	if floor != "" {
		params["floor"] = floor
	}
	if specialty != "" {
		params["specialty"] = specialty
	}
	if categoryID != "" {
		params["categoryId"] = categoryID
	}
	if family != "" {
		params["family"] = family
	}
	if familyType != "" {
		params["familyType"] = familyType
	}

	return o.GetIntgrElementsWithParams(fileID, params)
}
