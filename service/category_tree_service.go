// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/bean/common"
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

const (
	//获取文件转换的构件层次结构
	categoryURI string = "/data/hierarchy?fileId=%d"
)

//CategoryTreeService ***
type CategoryTreeService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewCategoryTreeService ***
func NewCategoryTreeService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *CategoryTreeService {
	o := &CategoryTreeService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *CategoryTreeService) categoryURL(fileId int64, isV2 bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+categoryURI, fileId)
	if isV2 {
		result = result + "&v=2.0"
	}
	return result
}

//-----------------------------------------------------------------------------------

//GetCategoryTreeResp 文件转换相关: 获取单文件的所有构件类别、族和族类型树
//http://static.bimface.com/book/restful/articles/api/translate/get-hierarchy.html
//1）获取1.0版本结果数据； 2）获取2.0版本结果数据
/***
字段	类型	必填	描述
fileId	Number	Y	文件Id
v		String	N	结果数据版本：1.0（1.0版本结果数据）2.0（2.0版本结果数据）	 默认为1.0数据
***/
func (o *CategoryTreeService) GetCategoryTreeResp(fileId int64, isV2 bool) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.categoryURL(fileId, isV2), headers.Header)
	return resp, nil
}

//GetCategoryTree 文件转换相关: 获取单文件的所有构件类别、族和族类型树, 结果数据版本：1.0（1.0版本结果数据）
/***
字段	类型	必填	描述
fileId	Number	Y	文件Id
***/
func (o *CategoryTreeService) GetCategoryTree(fileId int64) ([]*response.Category, error) {
	resp, err := o.GetCategoryTreeResp(fileId, false)
	if err != nil {
		return nil, err
	}

	result := make([]*response.Category, 0)
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//GetCategoryTreeV2 文件转换相关: 获取单文件的所有构件类别、族和族类型树, 结果数据版本：2.0（2.0版本结果数据）
/***
字段	类型	必填	描述
fileId	Number	Y	文件Id
***/
func (o *CategoryTreeService) GetCategoryTreeV2(fileId int64) ([]*common.TreeNode, error) {
	resp, err := o.GetCategoryTreeResp(fileId, true)
	if err != nil {
		return nil, err
	}

	result := make([]*common.TreeNode, 0)
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
