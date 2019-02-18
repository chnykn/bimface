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
	createDatabagByFileID      string = "/files/%d/offlineDatabag?callback=%s"
	createDatabagByIntegrateID string = "/integrations/%d/offlineDatabag?callback=%s"
	createDatabagByCompareID   string = "/comparisions/%d/offlineDatabag?callback=%s"

	queryDatabagByFileID      string = "/files/%d/offlineDatabag"
	queryDatabagByIntegrateID string = "/integrations/%d/offlineDatabag"
	queryDatabagByCompareID   string = "/comparisions/%d/offlineDatabag"

	getDatabagURLByFileID      string = "/data/databag/downloadUrl?fileId=%d&type=offline&databagVersion=%s"
	getDatabagURLByIntegrateID string = "/data/databag/downloadUrl?integrateId=%d&type=offline&databagVersion=%s"
	getDatabagURLByCompareID   string = "/data/databag/downloadUrl?comapreId=%d&type=offline&databagVersion=%s"
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
func (o *DatabagService) createDatabagURL(kind string, xxID int64, callback string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+createDatabagURI, kind, xxID)
	if callback != "" {
		result = result + "?callback=" + utils.EncodeURI(callback)
	}
	return result
}

func (o *DatabagService) queryDatabagURL(kind string, xxID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+queryDatabagURI, kind, xxID)
}

//databagVersion 数据包版本，如果只有一个，则下载唯一的数据包，如果多个，则必须指定数据包版本
func (o *DatabagService) downloadDatabagURL(kind string, xxID int64, databagVersion string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+getDatabagURI, kind, xxID)
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
fileId		Number	Y	通过文件转换ID创建离线数据包时必填
integrateId	Number	Y	通过集成模型ID创建离线数据包时必填
compareId	Number	Y	通过模型对比ID创建离线数据包时必填
callback	String	N	回调url
***/
func (o *DatabagService) CreateDatabag(databagRequest *request.DatabagRequest) (*response.Databag, error) {

	var url string
	if databagRequest.FileID != nil {
		url = o.createDatabagURL("files", *databagRequest.FileID, databagRequest.Callback)
	} else if databagRequest.IntegrateID != nil {
		url = o.createDatabagURL("integrations", *databagRequest.IntegrateID, databagRequest.Callback)
	} else if databagRequest.CompareID != nil {
		url = o.createDatabagURL("comparisions", *databagRequest.IntegrateID, databagRequest.Callback)
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
fileId		Number	Y	通过文件转换ID创建离线数据包时必填
integrateId	Number	Y	通过集成模型ID创建离线数据包时必填
compareId	Number	Y	通过模型对比ID创建离线数据包时必填
***/
func (o *DatabagService) QueryDatabag(databagRequest *request.DatabagRequest) ([]*response.Databag, error) {

	var url string
	if databagRequest.FileID != nil {
		url = o.queryDatabagURL("files", *databagRequest.FileID)
	} else if databagRequest.IntegrateID != nil {
		url = o.queryDatabagURL("integrations", *databagRequest.IntegrateID)
	} else if databagRequest.CompareID != nil {
		url = o.queryDatabagURL("comparisions", *databagRequest.IntegrateID)
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
fileId			Number	Y	通过文件转换ID获取离线数据包下载地址时必填
integrateId		Number	Y	通过集成模型ID获取离线数据包下载地址时必填
compareId		Number	Y	通过模型对比ID获取离线数据包下载地址时必填
type			String	Y	值必须是“offline”
databagVersion	String	N	数据包版本，如果只有一个，则下载唯一的数据包，如果多个，则必须指定数据包版本, 例如 3.0
***/
func (o *DatabagService) GetDatabagDownloadURL(databagRequest *request.DatabagRequest) (string, error) {

	var url string
	if databagRequest.FileID != nil {
		url = o.downloadDatabagURL("fileId", *databagRequest.FileID, databagRequest.DatabagVersion)
	} else if databagRequest.IntegrateID != nil {
		url = o.downloadDatabagURL("integrateId", *databagRequest.IntegrateID, databagRequest.DatabagVersion)
	} else if databagRequest.CompareID != nil {
		url = o.downloadDatabagURL("comapreId", *databagRequest.IntegrateID, databagRequest.DatabagVersion)
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
