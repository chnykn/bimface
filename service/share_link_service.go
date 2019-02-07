// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"strconv"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

const (
	//文件发起转换以后，可以根据fileId生成该文件的分享链接；
	//或者发起集成模型以后，可以根据integrateId生成集成模型的分享链接。
	createTranslateShareURI string = "/share?fileId=%d"      //&activeHours=%s
	createIntegrateShareURI string = "/share?integrateId=%d" //&activeHours=%s

	//若不希望继续分享，可以根据fileId或integrateId取消对应的分享链接，使之失效
	deleteTranslateShareURI string = "/share?fileId=%d"
	deleteIntegrateShareURI string = "/share?integrateId=%d"
)

//ShareLinkService ***
type ShareLinkService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewShareLinkService ***
func NewShareLinkService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *ShareLinkService {
	o := &ShareLinkService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *ShareLinkService) createTranslateShareURI(fileID int64, activeHours int) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+createTranslateShareURI, fileID)
	if activeHours > 0 {
		result = result + "&activeHours=" + strconv.Itoa(activeHours)
	}
	return result
}

func (o *ShareLinkService) deleteTranslateShareURL(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+deleteTranslateShareURI, fileID)
}

func (o *ShareLinkService) createIntegrateShareURI(integrateID int64, activeHours int) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+createIntegrateShareURI, integrateID)
	if activeHours > 0 {
		result = result + "&activeHours=" + strconv.Itoa(activeHours)
	}
	return result
}

func (o *ShareLinkService) deleteIntegrateShareURL(integrateID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+deleteIntegrateShareURI, integrateID)
}

//---------------------------------------------------------------------

func (o *ShareLinkService) generalCreateShare(isTranslate bool, xxID int64, activeHours int) (*response.ShareLink, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	var url string
	if isTranslate {
		url = o.createTranslateShareURI(xxID, activeHours)
	} else {
		url = o.createIntegrateShareURI(xxID, activeHours)
	}
	resp := o.ServiceClient.Post(url, headers.Header)

	result := response.NewShareLink("", "")
	err = utils.RespToBean(resp, result)

	return result, err
}

//CreateShare  文件发起转换以后，根据fileId生成该文件的分享链接
//http://static.bimface.com/book/restful/articles/api/share/create-sharelink.html
/***
字段		类型	必填	描述
fileId		Number	N 	(集成ID二选一)	文件ID
integrateId	Number	N 	(文件ID二选一)	集成ID
activeHours	Number	N	有效时长，单位：小时，如果不设置表示永久有效
***/
func (o *ShareLinkService) CreateShare(fileID int64, activeHours int) (*response.ShareLink, error) {
	return o.generalCreateShare(true, fileID, activeHours)
}

//CreateShareTranslation same to CreateShare
func (o *ShareLinkService) CreateShareTranslation(fileID int64, activeHours int) (*response.ShareLink, error) {
	return o.CreateShare(fileID, activeHours)
}

//CreateShareIntegration  发起集成模型以后，根据integrateId生成集成模型的分享链接
//http://static.bimface.com/book/restful/articles/api/share/create-sharelink.html
/***
字段		类型	必填	描述
fileId		Number	N 	(集成ID二选一)	文件ID
integrateId	Number	N 	(文件ID二选一)	集成ID
activeHours	Number	N	有效时长，单位：小时，如果不设置表示永久有效
***/
func (o *ShareLinkService) CreateShareIntegration(integrateID int64, activeHours int) (*response.ShareLink, error) {
	return o.generalCreateShare(false, integrateID, activeHours)
}

//---------------------------------------------------------------------

func (o *ShareLinkService) generalDeleteShare(isTranslate bool, xxID int64) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	var url string
	if isTranslate {
		url = o.deleteTranslateShareURL(xxID)
	} else {
		url = o.deleteIntegrateShareURL(xxID)
	}
	resp := o.ServiceClient.Delete(url, headers.Header)

	result, err := utils.RespToResult(resp)
	if err != nil {
		return result.Code, err
	}

	return result.Code, nil
}

//DeleteShare 取消分享: 文件转换以后 分享的链接
//http://static.bimface.com/book/restful/articles/api/share/delete-sharelink.html
func (o *ShareLinkService) DeleteShare(fileID int64) (string, error) {
	return o.generalDeleteShare(true, fileID)
}

//DeleteShareTranslation same to DeleteShare
func (o *ShareLinkService) DeleteShareTranslation(fileID int64) (string, error) {
	return o.DeleteShare(fileID)
}

//DeleteShareIntegration 取消分享: 集成模型以后 分享的链接
//http://static.bimface.com/book/restful/articles/api/share/delete-sharelink.html
func (o *ShareLinkService) DeleteShareIntegration(integrateID int64) (string, error) {
	return o.generalDeleteShare(false, integrateID)
}
