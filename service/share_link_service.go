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

	getSharesURI    string = "/shares"
	deleteSharesURI string = "/shares"
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
// expireDate's format YYYY-MM-DD
func (o *ShareLinkService) createTranslateShareURI(fileId int64, activeHours int, expireDate string, needPassword bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+createTranslateShareURI, fileId)
	if activeHours > 0 {
		result = result + "&activeHours=" + strconv.Itoa(activeHours)
	} else if expireDate != "" {
		result = result + "&expireDate=" + expireDate
	}

	if needPassword {
		result = result + "&needPassword=true"
	}
	return result
}

func (o *ShareLinkService) deleteTranslateShareURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+deleteTranslateShareURI, fileId)
}

// expireDate's format YYYY-MM-DD
func (o *ShareLinkService) createIntegrateShareURI(integrateId int64, activeHours int, expireDate string, needPassword bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+createIntegrateShareURI, integrateId)
	if activeHours > 0 {
		result = result + "&activeHours=" + strconv.Itoa(activeHours)
	} else if expireDate != "" {
		result = result + "&expireDate=" + expireDate
	}

	if needPassword {
		result = result + "&needPassword=true"
	}
	return result
}

func (o *ShareLinkService) deleteIntegrateShareURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+deleteIntegrateShareURI, integrateId)
}

//---------------------------------------------------------------------

func (o *ShareLinkService) generalCreateShare(isTranslate bool, xxId int64, activeHours int, expireDate string, needPassword bool) (*response.ShareLink, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	var url string
	if isTranslate {
		url = o.createTranslateShareURI(xxId, activeHours, expireDate, needPassword)
	} else {
		url = o.createIntegrateShareURI(xxId, activeHours, expireDate, needPassword)
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
fileId		Number	N 	(集成Id二选一)	文件Id
integrateId	Number	N 	(文件Id二选一)	集成Id
activeHours	Number	N	有效时长，单位：小时，如果不设置表示永久有效
***/
func (o *ShareLinkService) CreateShare(fileId int64, activeHours int, expireDate string, needPassword bool) (*response.ShareLink, error) {
	return o.generalCreateShare(true, fileId, activeHours, expireDate, needPassword)
}

//CreateShareTranslation same to CreateShare
func (o *ShareLinkService) CreateShareTranslation(fileId int64, activeHours int, expireDate string, needPassword bool) (*response.ShareLink, error) {
	return o.CreateShare(fileId, activeHours, expireDate, needPassword)
}

//CreateShareIntegration  发起集成模型以后，根据integrateId生成集成模型的分享链接
//http://static.bimface.com/book/restful/articles/api/share/create-sharelink.html
/***
字段		类型	必填	描述
fileId		Number	N 	(集成Id二选一)	文件Id
integrateId	Number	N 	(文件Id二选一)	集成Id
activeHours	Number	N	有效时长，单位：小时，如果不设置表示永久有效
***/
func (o *ShareLinkService) CreateShareIntegration(integrateId int64, activeHours int, expireDate string, needPassword bool) (*response.ShareLink, error) {
	return o.generalCreateShare(false, integrateId, activeHours, expireDate, needPassword)
}

//---------------------------------------------------------------------

func (o *ShareLinkService) generalDeleteShare(isTranslate bool, xxId int64) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	var url string
	if isTranslate {
		url = o.deleteTranslateShareURL(xxId)
	} else {
		url = o.deleteIntegrateShareURL(xxId)
	}
	resp := o.ServiceClient.Delete(url, headers.Header)

	result, err := utils.RespToResult(resp)
	if err != nil {
		if result != nil {
			return result.Code, err
		}
		return "", err
	}

	return result.Code, nil
}

//DeleteShare 取消分享: 文件转换以后 分享的链接
//http://static.bimface.com/book/restful/articles/api/share/delete-sharelink.html
func (o *ShareLinkService) DeleteShare(fileId int64) (string, error) {
	return o.generalDeleteShare(true, fileId)
}

//DeleteShareTranslation same to DeleteShare
func (o *ShareLinkService) DeleteShareTranslation(fileId int64) (string, error) {
	return o.DeleteShare(fileId)
}

//DeleteShareIntegration 取消分享: 集成模型以后 分享的链接
//http://static.bimface.com/book/restful/articles/api/share/delete-sharelink.html
func (o *ShareLinkService) DeleteShareIntegration(integrateId int64) (string, error) {
	return o.generalDeleteShare(false, integrateId)
}

//---------------------------------------------------------------------

func (o *ShareLinkService) GetShares(pageNo int, pageSize int) ([]*response.ShareLink, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	url := o.Endpoint.APIHost + getSharesURI
	if pageNo > 0 {
		url = url + "&pageNo=" + strconv.Itoa(pageNo)
	} else if pageSize > 0 {
		url = url + "&pageSize=" + strconv.Itoa(pageSize)
	}

	resp := o.ServiceClient.Get(url, headers.Header)

	result := make([]*response.ShareLink, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}

func (o *ShareLinkService) DeleteShares(soruceIds []int64) (string, error) {
	if len(soruceIds) <= 0 {
		return "", fmt.Errorf("sourceIds is null")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	//--------------

	url := o.Endpoint.APIHost + deleteSharesURI + "?sourceIds="

	sids := "["
	high := len(soruceIds) - 1
	for i := range soruceIds {
		sids = sids + strconv.FormatInt(soruceIds[i], 10)
		if i < high {
			sids = sids + ","
		}
	}
	sids = sids + "]"

	url = url + sids
	resp := o.ServiceClient.Delete(url, headers.Header)

	//--------------

	result, err := utils.RespToResult(resp)
	if err != nil {
		if result != nil {
			return result.Code, err
		}
		return "", err
	}

	return result.Code, nil
}
