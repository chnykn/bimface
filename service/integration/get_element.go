// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/httpkit"
)

const (
	//查询满足条件的构件ID列表 GET https://api.bimface.com/data/v2/integrations/{integrateId}/elementIds
	elementIdsURI string = "/data/v2/integrations/%d/elementIds"

	//获取指定构件的属性 GET https://api.bimface.com/data/v2/integrations/{integrateId}/elements/{elementId}
	elementPropertiesURI string = "/data/v2/integrations/%d/elements/%s"

	//获取子文件/链接内的指定构件的属性 GET https://api.bimface.com/data/v2/integrations/{integrateId}/files/{fileIdHash}/elements/{elementId}
	elementFilePropertiesURI string = "/data/v2/integrations/%d/files/%d/elements/%s"

	//获取多个构件的公共属性 POST https://api.bimface.com/data/v2/integrations/{integrateId}/commonElementProperties
	elementCommonPropertiesURI string = "/data/v2/integrations/%d/commonElementProperties"
)

//---------------------------------------------------------------------

func (o *Service) elementIdsURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+elementIdsURI, fileId)
}

func (o *Service) elementPropertiesURL(integrateId int64, elementId string, includeOverrides bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+elementPropertiesURI, integrateId, elementId)
	if includeOverrides {
		result = result + "?includeOverrides=true"
	}
	return result
}

func (o *Service) elementFilePropertiesURL(integrateId int64, fileId int64, elementId string, includeOverrides bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+elementFilePropertiesURI, integrateId, fileId, elementId)
	if includeOverrides {
		result = result + "?includeOverrides=true"
	}
	return result
}

func (o *Service) elementCommonPropertiesURL(integrateId int64, includeOverrides bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+elementCommonPropertiesURI, integrateId)

	if includeOverrides {
		result = result + "?includeOverrides=true"
	}
	return result
}

//-----------------------------------------------------------------------------------

// 查询满足条件的构件ID列表
// 必填参数: fileId  params相关参数，详见 https://static.bimface.com/restful-apidoc/dist/modelIntegration.html#_getelementidsusingget_1
func (o *Service) GetElementIdsWithParams(integrateId int64, params map[string]string) ([]string, error) {
	result := make([]string, 0)
	err := o.GET(o.elementIdsURL(integrateId), &result, httpkit.NewReqQuery(params))

	return result, err
}

//查询满足条件的构件ID列表
//https://static.bimface.com/restful-apidoc/dist/modelIntegration.html#_getelementidsusingget_1
/***
字段		类型	必填	描述
integrateId Number	Y	集成ID
specialty	String	N	专业
floor		String	N	楼层
categoryId	String	N	构件分类Id
roomId      String	N	房间id
family		String	N	族
familyType	String	N	族类型
systemType  String	N	系统类型
*** 其他参数详见网址地址
***/
func (o *Service) GetElementIds(integrateId int64, floor, specialty, categoryId,
	family, familyType string) ([]string, error) {

	params := make(map[string]string)
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

	return o.GetElementIdsWithParams(integrateId, params)
}

//------------------------------------------------------------------------------------

// 获取构件属性
func (o *Service) GetElementProperties(integrateId int64, elementId string, includeOverrides bool) (*response.PropertyBean, error) {
	result := new(response.PropertyBean)
	err := o.GET(o.elementPropertiesURL(integrateId, elementId, includeOverrides), result)

	return result, err
}

// 获取构件属性
func (o *Service) GetElementFileProperties(integrateId int64, fileId int64, elementId string, includeOverrides bool) (*response.PropertyBean, error) {
	result := new(response.PropertyBean)
	err := o.GET(o.elementFilePropertiesURL(integrateId, fileId, elementId, includeOverrides), result)

	return result, err
}

// 获取多个构件的共同属性
func (o *Service) GetElementCommonProperties(integrateId int64, elementIdsRequest *request.FileElementIdsRequest, includeOverrides bool) (*response.PropertyBean, error) {
	result := new(response.PropertyBean)

	body := httpkit.JsonReqBody(elementIdsRequest)
	err := o.POST(o.elementCommonPropertiesURL(integrateId, includeOverrides), result, body)

	return result, err
}
