// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

const (

	//文件bake数据包  https://api.bimface.com/files/{fileId}/bake
	fileBakeURI string = "/files/%d/bake"

	//集成模型bake数据包 https://api.bimface.com/integrations/{integrateId}/bake
	intgrBakeURI string = "/integrations/%d/bake"
)

//BakeService ***
type BakeService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewBakeService ***
func NewBakeService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *BakeService {
	o := &BakeService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *BakeService) fileBakeURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+fileBakeURI, fileId)
}

func (o *BakeService) intgrBakeURL(intgrId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+intgrBakeURI, intgrId)
}

//---------------------------------------------------------------------

//CreateFileBake 为文件创建bake数据包
//https://static.bimface.com/restful-apidoc/dist/bakeDatabag.html#_createtranslatebakedatabagusingput
func (o *BakeService) CreateFileBake(fileId int64, config map[string]string) (*response.Bake, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	conf := make(map[string]interface{})
	conf["config"] = config

	body := req.BodyJSON(conf)
	resp := o.ServiceClient.Put(o.fileBakeURL(fileId), headers.Header, body)

	result := response.NewBake()
	err = utils.RespToBean(resp, result)

	return result, err
}

//CreateFileDefaultBake ***
func (o *BakeService) CreateFileDefaultBake(fileId int64) (*response.Bake, error) {

	confg := make(map[string]string)
	confg["texture"] = "true" //texture:true

	return o.CreateFileBake(fileId, confg)
}

//GetFileBake 查询文件bake数据包
//https://static.bimface.com/restful-apidoc/dist/bakeDatabag.html#_gettranslatebakedatabagusingget
func (o *BakeService) GetFileBake(fileId int64) ([]*response.Bake, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.fileBakeURL(fileId), headers.Header)

	result := make([]*response.Bake, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}

//--------------------------------------------------------------------------------------------

//CreateIntgrBake 为集成模型创建bake数据包
//https://static.bimface.com/restful-apidoc/dist/bakeDatabag.html#_createintegratebakedatabagusingput
func (o *BakeService) CreateIntgrBake(intgrId int64, config map[string]string) (*response.Bake, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	conf := make(map[string]interface{})
	conf["config"] = config

	body := req.BodyJSON(conf)
	resp := o.ServiceClient.Put(o.intgrBakeURL(intgrId), headers.Header, body)

	result := response.NewBake()
	err = utils.RespToBean(resp, result)

	return result, err
}

//CreateIntgrDefaultBake ***
func (o *BakeService) CreateIntgrDefaultBake(intgrId int64) (*response.Bake, error) {

	confg := make(map[string]string)
	confg["texture"] = "true" //texture:true

	return o.CreateIntgrBake(intgrId, confg)
}

//GetIntgrBake 查询集成模型bake数据包
//https://static.bimface.com/restful-apidoc/dist/bakeDatabag.html#_getintegratebakedatabagusingget
func (o *BakeService) GetIntgrBake(intgrId int64) ([]*response.Bake, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Put(o.intgrBakeURL(intgrId), headers.Header)

	result := make([]*response.Bake, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}
