// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/v2/bean/request"
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	translateURI        string = "/translate"
	getTranslateURI     string = "/translate?fileId=%d"
	translateDetailsURI string = "/translateDetails"
)

//---------------------------------------------------------------------

func (o *Service) translateURL() string {
	return o.Endpoint.APIHost + translateURI
}

func (o *Service) getTranslateURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getTranslateURI, fileId)
}

func (o *Service) translateDetailsURL() string {
	return fmt.Sprintf(o.Endpoint.APIHost + translateDetailsURI)
}

//-----------------------------------------------------------------------------------

//发起文件转换
func (o *Service) Translate(transRequest *request.FileTranslateRequest) (*response.FileTranslateBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := req.BodyJSON(transRequest)
	resp := o.ServiceClient.Put(o.translateURL(), body, headers.Header)

	var result *response.FileTranslateBean
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//获取转换状态
func (o *Service) GetStatus(fileId int64) (*response.FileTranslateBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getTranslateURL(fileId), headers.Header)

	var result *response.FileTranslateBean
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//批量获取转换状态详情
func (o *Service) GetDetails(queryRequest *request.TranslateQueryRequest) (*response.FileTranslateDetailBeanPageList, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	var resp *req.Resp
	if queryRequest != nil {
		body := req.BodyJSON(queryRequest)
		resp = o.ServiceClient.Post(o.translateDetailsURL(), body, headers.Header)
	} else {
		resp = o.ServiceClient.Post(o.translateDetailsURL(), headers.Header)
	}

	var result *response.FileTranslateDetailBeanPageList
	err = utils.RespToBean(resp, result)

	return result, err
}
