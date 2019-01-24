// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chnykn/bimface/bean/request"
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/http"
	"github.com/chnykn/bimface/utils"
	"fmt"

	"github.com/imroc/req"
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
func NewCompareService(serviceClient *http.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *CompareService {
	o := &CompareService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //http.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *CompareService) compareURL() string {
	return o.Endpoint.APIHost + compareURI
}

func (o *CompareService) getCompareURL(compareID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getCompareURI, compareID)
}

func (o *CompareService) compareDataURL(compareID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+compareDataURI, compareID)
}

/*
previousFileId		Number	Y	对比差异构件来源文件ID
previousElementId	String	Y	对比差异构件来源构件ID
followingFileId		Number	Y	对比差异构件变更文件ID
followingElementId	String	Y	对比差异构件互为变更构件ID
*/
func (o *CompareService) compareElementDataURL(compareID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+compareElementDataURI, compareID)
}

//---------------------------------------------------------------------

//Compare 发起模型对比
//http://static.bimface.com/book/restful/articles/api/compare/post-compare.html
/***
字段			类型		必填	描述
previousFileId	Number		N	变更前文件ID，如果为新增文件，则为null
followingFileId	Number		N	变更后文件ID，如果为删除文件，则为null
sources			Source[]	Y	数组，多个CompareSource，待对比的文件
name			String		N	用户指定对比后的模型的名字
sourceId		String		N	第三方应用自己的ID
priority		Number		N	优先级，数字越大，优先级越低	1, 2, 3
callback		String		N	Callback地址，待对比完毕以后，BIMFace会回调该地址
***/
func (o *CompareService) Compare(compareRequst *request.CompareRequest) (*response.CompareStatus, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := req.BodyJSON(compareRequst)
	resp := o.ServiceClient.Post(o.compareURL(), headers.Header, body)

	result := response.NewCompareStatus()
	err = http.RespToBean(resp, result)

	return result, err
}

//GetCompare 获取模型对比状态
//http://static.bimface.com/book/restful/articles/api/compare/get-compare.html
/***
字段		类型	必填	描述
compareId	Number	Y	模型对比ID
***/
func (o *CompareService) GetCompare(compareID int64) (*response.CompareStatus, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getCompareURL(compareID), headers.Header)

	result := response.NewCompareStatus()
	err = http.RespToBean(resp, result)

	return result, err
}

//GetCompareStatus same to GetCompare
func (o *CompareService) GetCompareStatus(compareID int64) (*response.CompareStatus, *utils.Error) {
	return o.GetCompare(compareID)
}

//GetCompareData 获取模型对比结果
//http://static.bimface.com/book/restful/articles/api/compare/get-compare-rst.html
/***
字段		类型	必填	描述
compareId	Number	Y	模型对比Id
***/
func (o *CompareService) GetCompareData(compareID int64) ([]response.CompareData, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.compareDataURL(compareID), headers.Header)

	result, err := http.RespToBeans(resp, &response.CompareData{})
	return result.([]response.CompareData), nil
}

//GetCompareElementDiffWithParams 获取修改构件属性差异
//http://static.bimface.com/book/restful/articles/api/compare/get-compare-ele-diff.html
func (o *CompareService) GetCompareElementDiffWithParams(compareID int64, params req.QueryParam) (*response.ElementDiff, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.compareElementDataURL(compareID), params, headers.Header)

	result := response.NewElementDiff()
	err = http.RespToBean(resp, result)

	return result, err
}

//GetCompareElementDiff 获取修改构件属性差异, same to GetCompareElementDiffWithParams
/***
字段				类型	必填	描述
compareId			Number	Y	模型对比ID
previousFileId		Number	Y	对比差异构件来源文件ID
previousElementId	String	Y	对比差异构件来源构件ID
followingFileId		Number	Y	对比差异构件变更文件ID
followingElementId	String	Y	对比差异构件互为变更构件ID
***/
func (o *CompareService) GetCompareElementDiff(compareID int64, previousFileID int64, previousElementID string,
	followingFileID int64, followingElementID string) (*response.ElementDiff, *utils.Error) {

	params := make(req.QueryParam)

	if previousFileID > 0 {
		params["previousFileId"] = previousFileID
	}
	if previousElementID != "" {
		params["previousElementId"] = previousElementID
	}
	if followingFileID > 0 {
		params["followingFileId"] = followingFileID
	}
	if followingElementID != "" {
		params["followingElementId"] = followingElementID
	}

	return o.GetCompareElementDiffWithParams(compareID, params)
}
