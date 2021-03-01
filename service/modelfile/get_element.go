// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"
	"strings"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/v2/bean/request"
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	//查询满足条件的构件ID列表 GET https://api.bimface.com/data/v2/files/{fileId}/elementIds
	elementIdsURI string = "/data/v2/files/%d/elementIds"

	//获取构件属性 GET https://api.bimface.com/data/v2/files/{fileId}/elements/{elementId}
	elementPropertiesURI string = "/data/v2/files/%d/elements/%s"

	//获取多个构件的共同属性 GET https://api.bimface.com/data/v2/files/{fileId}/commonElementProperties
	elementCommonPropertiesURI string = "/data/v2/files/%d/commonElementProperties?elementIds=%s"

	//批量获取构件属性 POST https://api.bimface.com/data/v2/files/{fileId}/elements
	elementsPropertiesURI string = "/data/v2/files/%d/elements"
)

//---------------------------------------------------------------------

func (o *Service) elementIdsURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+elementIdsURI, fileId)
}

func (o *Service) elementPropertiesURL(fileId int64, elementId string, includeOverrides bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+elementPropertiesURI, fileId, elementId)
	if includeOverrides {
		result = result + "?includeOverrides=true"
	}
	return result
}

func (o *Service) elementCommonPropertiesURL(fileId int64, elementIds []string, includeOverrides bool) string {
	elemIds := ""
	if len(elementIds) > 0 {
		for _, s := range elementIds {
			elemIds = elemIds + s + ","
		}
		elemIds = strings.TrimRight(elemIds, ",")
	}

	result := fmt.Sprintf(o.Endpoint.APIHost+elementCommonPropertiesURI, fileId, elemIds)

	if includeOverrides {
		result = result + "?includeOverrides=true"
	}
	return result
}

func (o *Service) elementsPropertiesURL(fileId int64, includeOverrides bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+elementsPropertiesURI, fileId)

	if includeOverrides {
		result = result + "?includeOverrides=true"
	}
	return result
}

//-----------------------------------------------------------------------------------

//查询满足条件的构件ID列表
//必填参数: fileId  params相关参数，详见 http://static.bimface.com/restful-apidoc/dist/translateSingleModel.html#_getelementidsusingget
func (o *Service) GetElementIdsWithParams(fileId int64, params req.QueryParam) ([]string, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.elementIdsURL(fileId), params, headers.Header)

	result := make([]string, 0)
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//查询满足条件的构件ID列表
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
func (o *Service) GetElementIds(fileId int64, floor, specialty, categoryId,
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

	return o.GetElementIdsWithParams(fileId, params)
}

//------------------------------------------------------------------------------------

//获取构件属性
func (o *Service) GetElementProperties(fileId int64, elementId string, includeOverrides bool) (*response.PropertyBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.elementPropertiesURL(fileId, elementId, includeOverrides), headers.Header)

	result := new(response.PropertyBean)
	err = utils.RespToBean(resp, result)

	return result, err
}

//获取多个构件的共同属性
func (o *Service) GetElementCommonProperties(fileId int64, elementIds []string, includeOverrides bool) (*response.PropertyBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.elementCommonPropertiesURL(fileId, elementIds, includeOverrides), headers.Header)

	result := new(response.PropertyBean)
	err = utils.RespToBean(resp, result)

	return result, err
}

//批量获取构件属性
func (o *Service) GetElementsProperties(fileId int64, filterRequest *request.PropertyFilterRequest, includeOverrides bool) ([]*response.PropertyBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := req.BodyJSON(filterRequest)
	resp := o.ServiceClient.Post(o.elementsPropertiesURL(fileId, includeOverrides), body, headers.Header)

	result := make([]*response.PropertyBean, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}
