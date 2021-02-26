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

func (o *ElementService) intgrElementURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+intgrElementURI, integrateId)
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
