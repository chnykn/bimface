// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/chnykn/bimface/bean/request"
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"

	"github.com/imroc/req"
)

const (
	translateURI    string = "/translate"
	getTranslateURI string = "/translate?fileId=%d"
)

//TranslateService ***
type TranslateService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewTranslateService ***
func NewTranslateService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *TranslateService {
	o := &TranslateService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *TranslateService) translateURL() string {
	return o.Endpoint.APIHost + translateURI
}

func (o *TranslateService) getTranslateURL(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getTranslateURI, fileID)
}

//-----------------------------------------------------------------------------------

//Translate 文件转换相关: 发起文件转换
//http://static.bimface.com/book/restful/articles/api/translate/put-translate.html
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
func (o *TranslateService) Translate(transRequest *request.TranslateRequest) (*response.TranslateStatus, error) {
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

//GetTranslateStatusResp ***
func (o *TranslateService) GetTranslateStatusResp(fileID int64) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getTranslateURL(fileID), headers.Header)
	return resp, err
}

//GetTranslateStatus 文件转换相关: 获取转换状态
//http://static.bimface.com/book/restful/articles/api/translate/get-translate.html
func (o *TranslateService) GetTranslateStatus(fileID int64) (*response.TranslateStatus, error) {
	resp, err := o.GetTranslateStatusResp(fileID)
	if err != nil {
		return nil, err
	}

	result := response.NewTranslateStatus()
	err = utils.RespToBean(resp, result)

	return result, err
}
