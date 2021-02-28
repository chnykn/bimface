// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sharelink

import (
	"fmt"
	"strconv"

	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	//分享列表 GET https://api.bimface.com/shares
	shareLinksURI string = "/shares"
)

func (o *Service) getShareLinkURL(isFile bool, objectId int64) string {

	var result string
	if isFile {
		result = fmt.Sprintf(o.Endpoint.APIHost+fileShareLinkURI, objectId)
	} else {
		result = fmt.Sprintf(o.Endpoint.APIHost+integrateShareLinkURI, objectId)
	}

	return result
}

func (o *Service) getShareLinksURL(pageNo int, pageSize int) string {

	if pageNo <= 0 {
		pageNo = 1
	}

	if pageSize <= 0 {
		pageSize = 20
	}

	return o.Endpoint.APIHost + shareLinksURI +
		"?pageNo=" + strconv.Itoa(pageNo) +
		"&pageSize=" + strconv.Itoa(pageSize)
}

//----------------------------------------------

func (o *Service) doGetShareLink(isFile bool, objectId int64) (*response.ShareLinkBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getShareLinkURL(isFile, objectId), headers.Header)

	var result *response.ShareLinkBean
	err = utils.RespToBean(resp, result)

	return result, err
}

//获取模型文件的分享链接
func (o *Service) GetFileShareLink(fileId int64) (*response.ShareLinkBean, error) {
	return o.doGetShareLink(true, fileId)
}

//获取模型集成的分享链接
func (o *Service) GetIntegrateShareLink(integrateId int64) (*response.ShareLinkBean, error) {
	return o.doGetShareLink(false, integrateId)
}

//--------------------------------------------------------------------------

//分享链接列表
func (o *Service) GetShareLinks(pageNo int, pageSize int) (*response.ShareLinkBeanPageList, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getShareLinksURL(pageNo, pageSize), headers.Header)

	var result *response.ShareLinkBeanPageList
	err = utils.RespToBean(resp, result)

	return result, err
}
