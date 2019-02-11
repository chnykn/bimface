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
	//发起模型集成
	integrateURI string = "/integrate"

	//获取集成状态
	getIntegrateURI string = "/integrate?integrateId=%d"

	//删除集成模型
	deleteIntegrateURI string = "/integrate?integrateId=%d"
)

//IntegrateService ***
type IntegrateService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewIntegrateService ***
func NewIntegrateService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *IntegrateService {
	o := &IntegrateService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *IntegrateService) integrateURL() string {
	return o.Endpoint.APIHost + integrateURI
}

func (o *IntegrateService) getIntegrateURL(integrateID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getIntegrateURI, integrateID)
}

func (o *IntegrateService) deleteIntegrateURL(integrateID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+deleteIntegrateURI, integrateID)
}

//-----------------------------------------------------------------------------------

//Integrate 模型集成相关: 发起模型集成
//http://doc.bimface.com/book/restful/articles/api/integrate/put-integrate.html
/***
字段					类型		必填	描述	示例
sources					Source[]	Y	待集成的文件列表
source.fileId			Number		Y	待集成的源文件ID，必须是 rvt 文件
source.specialty		String		N	待集成源文件对应的专业名称	AR
source.specialtySort	Number		N	显示专业层次结构时，排序数值越小，排序越前
source.floor			String		N	待集成源文件对应的楼层名称	F01
source.floorSort		Number		N	显示楼层层次结构时，排序数值越小，排序越前
sourceId				String		N	调用方的源ID
name					String		N	调用方设置的名称
priority				Number		Y	优先级，数字越大，优先级越低	1, 2, 3
callback				String		N	Callback地址，待集成完毕以后，BIMFACE会回调该地址
***/
func (o *IntegrateService) Integrate(intgrRequest *request.IntgrRequest) (*response.IntgrStatus, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := req.BodyJSON(intgrRequest)
	resp := o.ServiceClient.Put(o.integrateURL(), headers.Header, body)

	result := response.NewIntgrStatus()
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//GetIntegrateStatusResp ***
func (o *IntegrateService) GetIntegrateStatusResp(integrateID int64) (*req.Resp, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getIntegrateURL(integrateID), headers.Header)
	return resp, err
}

//GetIntegrateStatus 模型集成相关: 获取集成状态
//http://doc.bimface.com/book/restful/articles/api/integrate/get-integrate.html
func (o *IntegrateService) GetIntegrateStatus(integrateID int64) (*response.IntgrStatus, error) {
	resp, err := o.GetIntegrateStatusResp(integrateID)
	if err != nil {
		return nil, err
	}

	result := response.NewIntgrStatus()
	err = utils.RespToBean(resp, result)

	return result, err
}

//-----------------------------------------------------------------------------------

//DeleteIntegrate 模型集成相关: 删除集成模型
//http://doc.bimface.com/book/restful/articles/api/integrate/delete-integrate.html
func (o *IntegrateService) DeleteIntegrate(integrateID int64) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Delete(o.deleteIntegrateURL(integrateID), headers.Header)

	result, err := utils.RespToResult(resp)
	if err != nil {
		return result.Code, err
	}

	return result.Code, nil
}
