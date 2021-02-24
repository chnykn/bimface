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
)

/**
const (
	createDatabagByFileId      string = "/files/%d/offlineDatabag?callback=%s"
	createDatabagByIntegrateId string = "/integrations/%d/offlineDatabag?callback=%s"
	createDatabagByCompareId   string = "/comparisions/%d/offlineDatabag?callback=%s"

	queryDatabagByFileId      string = "/files/%d/offlineDatabag"
	queryDatabagByIntegrateId string = "/integrations/%d/offlineDatabag"
	queryDatabagByCompareId   string = "/comparisions/%d/offlineDatabag"

	getDatabagURLByFileId      string = "/data/databag/downloadUrl?fileId=%d&type=offline&databagVersion=%s"
	getDatabagURLByIntegrateId string = "/data/databag/downloadUrl?integrateId=%d&type=offline&databagVersion=%s"
	getDatabagURLByCompareId   string = "/data/databag/downloadUrl?comapreId=%d&type=offline&databagVersion=%s"
)
**/

const (
	createDatabagURI string = "/%s/%d/offlineDatabag" //?callback=%s
	queryDatabagURI  string = "/%s/%d/offlineDatabag"
	getDatabagURI    string = "/data/databag/downloadUrl?%s=%d&type=offline" //&databagVersion=%s
)

//DatabagService ***
type DatabagService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewDatabagService ***
func NewDatabagService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *DatabagService {
	o := &DatabagService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------
// kind must in [files, integrations, comparisions]
func (o *DatabagService) createDatabagURL(kind string, xxId int64, callback string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+createDatabagURI, kind, xxId)
	if callback != "" {
		result = result + "?callback=" + utils.EncodeURI(callback)
	}
	return result
}

func (o *DatabagService) queryDatabagURL(kind string, xxId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+queryDatabagURI, kind, xxId)
}

//databagVersion 数据包版本，如果只有一个，则下载唯一的数据包，如果多个，则必须指定数据包版本
func (o *DatabagService) downloadDatabagURL(kind string, xxId int64, databagVersion string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+getDatabagURI, kind, xxId)
	if databagVersion != "" {
		result = result + "&databagVersion=" + databagVersion
	}
	return result
}

//-------------------------------------------------------------------------------

//CreateDatabag 离线数据包相关: 创建离线数据包
//http://static.bimface.com/book/restful/articles/api/offlinedatabag/create-offlinedatabag.html
/***
字段		类型	必填	描述
fileId		Number	Y	通过文件转换Id创建离线数据包时必填
integrateId	Number	Y	通过集成模型Id创建离线数据包时必填
compareId	Number	Y	通过模型对比Id创建离线数据包时必填
callback	String	N	回调url
***/
func (o *DatabagService) CreateDatabag(databagRequest *request.DatabagRequest) (*response.Databag, error) {

	var url string
	if databagRequest.FileId != nil {
		url = o.createDatabagURL("files", *databagRequest.FileId, databagRequest.Callback)
	} else if databagRequest.IntegrateId != nil {
		url = o.createDatabagURL("integrations", *databagRequest.IntegrateId, databagRequest.Callback)
	} else if databagRequest.CompareId != nil {
		url = o.createDatabagURL("comparisions", *databagRequest.IntegrateId, databagRequest.Callback)
	}
	if url == "" {
		return nil, fmt.Errorf("url is null @ DatabagService.createDatabag")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Put(url, headers.Header)

	result := response.NewDatabag()
	err = utils.RespToBean(resp, result)

	return result, err
}

//-------------------------------------------------------------------------------

//QueryDatabag 离线数据包相关: 查询离线数据包
//http://static.bimface.com/book/restful/articles/api/offlinedatabag/query-offlinedataba.html
/***
字段		类型	必填	描述
fileId		Number	Y	通过文件转换Id创建离线数据包时必填
integrateId	Number	Y	通过集成模型Id创建离线数据包时必填
compareId	Number	Y	通过模型对比Id创建离线数据包时必填
***/
func (o *DatabagService) QueryDatabag(databagRequest *request.DatabagRequest) ([]*response.Databag, error) {

	var url string
	if databagRequest.FileId != nil {
		url = o.queryDatabagURL("files", *databagRequest.FileId)
	} else if databagRequest.IntegrateId != nil {
		url = o.queryDatabagURL("integrations", *databagRequest.IntegrateId)
	} else if databagRequest.CompareId != nil {
		url = o.queryDatabagURL("comparisions", *databagRequest.IntegrateId)
	}
	if url == "" {
		return nil, fmt.Errorf("url is null @ DatabagService.queryDatabag")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(url, headers.Header)

	result := make([]*response.Databag, 0)
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//-------------------------------------------------------------------------------

//GetDatabagDownloadURL 离线数据包相关: 获取离线数据包下载地址
//http://static.bimface.com/book/restful/articles/api/offlinedatabag/get-download-offlinedataba-url.html
/***
字段			类型	必填	描述
fileId			Number	Y	通过文件转换Id获取离线数据包下载地址时必填
integrateId		Number	Y	通过集成模型Id获取离线数据包下载地址时必填
compareId		Number	Y	通过模型对比Id获取离线数据包下载地址时必填
type			String	Y	值必须是“offline”
databagVersion	String	N	数据包版本，如果只有一个，则下载唯一的数据包，如果多个，则必须指定数据包版本, 例如 3.0
***/
func (o *DatabagService) GetDatabagDownloadURL(databagRequest *request.DatabagRequest) (string, error) {

	var url string
	if databagRequest.FileId != nil {
		url = o.downloadDatabagURL("fileId", *databagRequest.FileId, databagRequest.DatabagVersion)
	} else if databagRequest.IntegrateId != nil {
		url = o.downloadDatabagURL("integrateId", *databagRequest.IntegrateId, databagRequest.DatabagVersion)
	} else if databagRequest.CompareId != nil {
		url = o.downloadDatabagURL("comapreId", *databagRequest.IntegrateId, databagRequest.DatabagVersion)
	}
	if url == "" {
		return "", fmt.Errorf("url is null @ DatabagService.queryDatabag")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(url, headers.Header)

	result := new(string)
	err = utils.RespToBean(resp, result)

	if err != nil {
		return "", err
	}

	return *result, nil
}
