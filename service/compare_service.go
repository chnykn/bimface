// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/bean/request"
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

const (
	compareURI            string = "/compare"
	getCompareURI         string = "/compare?compareId=%d"
	compareDataURI        string = "/data/compare?compareId=%d"
	compareElementDataURI string = "data/compare/element/?compareId=%d"
)

//CompareService ***
type CompareService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewCompareService ***
func NewCompareService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *CompareService {
	o := &CompareService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *CompareService) compareURL() string {
	return o.Endpoint.APIHost + compareURI
}

func (o *CompareService) getCompareURL(compareId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getCompareURI, compareId)
}

func (o *CompareService) compareDataURL(compareId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+compareDataURI, compareId)
}

/*
previousFileId		Number	Y	对比差异构件来源文件Id
previousElementId	String	Y	对比差异构件来源构件Id
followingFileId		Number	Y	对比差异构件变更文件Id
followingElementId	String	Y	对比差异构件互为变更构件Id
*/
func (o *CompareService) compareElementDataURL(compareId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+compareElementDataURI, compareId)
}

//-----------------------------------------------------------------------------------

//Compare 发起模型对比
//http://static.bimface.com/book/restful/articles/api/compare/post-compare.html
/***
字段			类型		必填	描述
previousFileId	Number		N	变更前文件Id，如果为新增文件，则为null
followingFileId	Number		N	变更后文件Id，如果为删除文件，则为null
sources			Source[]	Y	数组，多个CompareSource，待对比的文件
name			String		N	用户指定对比后的模型的名字
sourceId		String		N	第三方应用自己的Id
priority		Number		N	优先级，数字越大，优先级越低	1, 2, 3
callback		String		N	Callback地址，待对比完毕以后，BIMFace会回调该地址
***/
func (o *CompareService) Compare(compareRequst *request.CompareRequest) (*response.CompareStatus, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := req.BodyJSON(compareRequst)
	resp := o.ServiceClient.Post(o.compareURL(), headers.Header, body)

	result := response.NewCompareStatus()
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//GetCompareStatusResp ***
func (o *CompareService) GetCompareStatusResp(compareId int64) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getCompareURL(compareId), headers.Header)
	return resp, err
}

//GetCompareStatus 获取模型对比状态
//http://static.bimface.com/book/restful/articles/api/compare/get-compare.html
/***
字段		类型	必填	描述
compareId	Number	Y	模型对比Id
***/
func (o *CompareService) GetCompareStatus(compareId int64) (*response.CompareStatus, error) {
	resp, err := o.GetCompareStatusResp(compareId)
	if err != nil {
		return nil, err
	}

	result := response.NewCompareStatus()
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//GetCompareDataResp ***
func (o *CompareService) GetCompareDataResp(compareId int64) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.compareDataURL(compareId), headers.Header)
	return resp, nil
}

//GetCompareData 获取模型对比结果
//http://static.bimface.com/book/restful/articles/api/compare/get-compare-rst.html
/***
字段		类型	必填	描述
compareId	Number	Y	模型对比Id
***/
func (o *CompareService) GetCompareData(compareId int64) ([]*response.CompareData, error) {
	resp, err := o.GetCompareDataResp(compareId)
	if err != nil {
		return nil, err
	}

	result := make([]*response.CompareData, 0)
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//-----------------------------------------------------------------------------------

//GetCompareElementResp ***
func (o *CompareService) GetCompareElementResp(compareId int64, params req.QueryParam) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.compareElementDataURL(compareId), params, headers.Header)
	return resp, err
}

//GetCompareElementDiffWithParams 获取修改构件属性差异
//http://static.bimface.com/book/restful/articles/api/compare/get-compare-ele-diff.html
func (o *CompareService) GetCompareElementDiffWithParams(compareId int64, params req.QueryParam) (*response.ElementDiff, error) {
	resp, err := o.GetCompareElementResp(compareId, params)
	if err != nil {
		return nil, err
	}

	result := response.NewElementDiff()
	err = utils.RespToBean(resp, result)

	return result, err
}

//GetCompareElementDiff 获取修改构件属性差异, same to GetCompareElementDiffWithParams
/***
字段				类型	必填	描述
compareId			Number	Y	模型对比Id
previousFileId		Number	Y	对比差异构件来源文件Id
previousElementId	String	Y	对比差异构件来源构件Id
followingFileId		Number	Y	对比差异构件变更文件Id
followingElementId	String	Y	对比差异构件互为变更构件Id
***/
func (o *CompareService) GetCompareElementDiff(compareId int64, previousFileId int64, previousElementId string,
	followingFileId int64, followingElementId string) (*response.ElementDiff, error) {

	params := make(req.QueryParam)

	if previousFileId > 0 {
		params["previousFileId"] = previousFileId
	}
	if previousElementId != "" {
		params["previousElementId"] = previousElementId
	}
	if followingFileId > 0 {
		params["followingFileId"] = followingFileId
	}
	if followingElementId != "" {
		params["followingElementId"] = followingElementId
	}

	return o.GetCompareElementDiffWithParams(compareId, params)
}
