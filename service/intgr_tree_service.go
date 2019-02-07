// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

//"Intgr" stand for "Integration"

import (
	"fmt"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"

	"github.com/imroc/req"
)

const (

	//获取集成模型的构件层次结构
	integrationTreeURI string = "/data/integration/tree?integrateId=%d&treeType=%d"
)

//IntgrTreeService ***
type IntgrTreeService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewIntgrTreeService ***
func NewIntgrTreeService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *IntgrTreeService {
	o := &IntgrTreeService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *IntgrTreeService) integrationTreeURL(integrateID int64, treeType int) string {
	return fmt.Sprintf(o.Endpoint.APIHost+integrationTreeURI, integrateID, treeType)
}

//-----------------------------------------------------------------------------------

//GetTreeResp 模型集成相关: 获取集成模型的构件层次结构
//模型集成以后，可以获取两种构件的层次结构：1）按专业视图；2）按楼层视图
//http://static.bimface.com/book/restful/articles/api/integrate/get-integrate-tree.html
/***
字段		类型	必填	描述
integrateId	Number	Y	集成ID
treeType	Number	Y	树类型：1（按专业视图）2（按楼层视图）
***/
func (o *IntgrTreeService) GetTreeResp(integrateID int64, treeType int) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integrationTreeURL(integrateID, treeType), headers.Header)
	return resp, nil
}

//GetSpecialtyTree 按专业视图,获取集成模型的构件层次结构
/***
字段		类型	必填	描述
integrateId	Number	Y	集成ID
***/
func (o *IntgrTreeService) GetSpecialtyTree(integrateID int64) (*response.IntgrSpecialtyTree, error) {
	resp, err := o.GetTreeResp(integrateID, 1)
	if err != nil {
		return nil, err
	}

	result := response.NewIntgrSpecialtyTree(1)
	err = utils.RespToBean(resp, result)

	return result, err
}

//GetFloorTree 按楼层视图,获取集成模型的构件层次结构
/***
字段		类型	必填	描述
integrateId	Number	Y	集成ID
***/
func (o *IntgrTreeService) GetFloorTree(integrateID int64) (*response.IntgrFloorTree, error) {
	resp, err := o.GetTreeResp(integrateID, 2)
	if err != nil {
		return nil, err
	}

	result := response.NewIntgrFloorTree(2)
	err = utils.RespToBean(resp, result)

	return result, err
}
