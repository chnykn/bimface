// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/bean/request"
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
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

//Translate 文件转换相关: 发起文件转换
/***
字段		类型		必填	描述
fileId		Number		Y	文件Id，即调用上传文件API返回的fileId
compressed	Boolean		N	是否为压缩文件，默认为false
rootName	String		N	如果是压缩文件，必须指定压缩包中哪一个是主文件，例如 root.rvt
priority	Number		优先级，数字越大，优先级越低	1, 2, 3
callback	String		N	Callback地址，待转换完毕以后，BIMFace会回调该地址
config		Json Object	N	转换引擎自定义参数，config参数跟转换引擎相关，不同的转换引擎支持不同的config格式。
							例如转换时添加内置材质，则添加参数值{“texture”:true}，
							添加外部材质时参考“使用模型外置材质场景”请求报文{“texture”:true}等
***/
func (o *Service) Translate(transRequest *request.TranslateRequest) (*response.TranslateStatus, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := req.BodyJSON(transRequest)
	resp := o.ServiceClient.Put(o.translateURL(), body, headers.Header)

	result := response.NewTranslateStatus()
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//TranslateStatusResp ***
func (o *Service) GetTranslateStatusResp(fileId int64) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getTranslateURL(fileId), headers.Header)
	return resp, err
}

//TranslateStatus 文件转换相关: 获取转换状态
func (o *Service) GetTranslateStatus(fileId int64) (*response.TranslateStatus, error) {
	resp, err := o.GetTranslateStatusResp(fileId)
	if err != nil {
		return nil, err
	}

	result := response.NewTranslateStatus()
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//TranslateDetailStatusResp ***
func (o *Service) GetTranslateDetailStatusResp(detailRequst *request.TranslateDetailRequest) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := req.BodyJSON(detailRequst)
	resp := o.ServiceClient.Post(o.translateDetailsURL(), body, headers.Header)

	return resp, err
}

//TranslateDetailStatus 批量获取转换状态详情
func (o *Service) GetTranslateDetailStatus(detailRequst *request.TranslateDetailRequest) (*response.TranslateDetailStatus, error) {
	resp, err := o.GetTranslateDetailStatusResp(detailRequst)
	if err != nil {
		return nil, err
	}

	result := response.NewTranslateDetailStatus()
	err = utils.RespToBean(resp, result)

	return result, err
}
