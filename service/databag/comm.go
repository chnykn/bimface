// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package databag

import (
	"fmt"

	"github.com/chnykn/bimface/v2/bean/request"
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	dataBagURI         string = "/%s/%d/offlineDatabag"
	downloadDataBagURI string = "/data/databag/downloadUrl?%s=%d&type=offline" //&databagVersion=%s
)

//---------------------------------------------------------------------
// kind must in [files, integrations, comparisions]
func (o *Service) dataBagURL(kind string, objectId int64, callback string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+dataBagURI, kind, objectId)
	if callback != "" {
		result = result + "?callback=" + utils.EncodeURI(callback)
	}
	return result
}

//dataBagVersion : 数据包版本，如果只有一个，则下载唯一的数据包，如果多个，则必须指定数据包版本
func (o *Service) downloadDataBagURL(kind string, objectId int64, dataBagVersion string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+downloadDataBagURI, kind, objectId)
	if dataBagVersion != "" {
		result = result + "&databagVersion=" + dataBagVersion
	}
	return result
}

//-------------------------------------------------------------------------------

//创建离线数据包
/***
字段		类型	必填	描述
fileId		Number	Y	通过文件转换Id创建离线数据包时必填
integrateId	Number	Y	通过集成模型Id创建离线数据包时必填
compareId	Number	Y	通过模型对比Id创建离线数据包时必填
callback	String	N	回调url
***/
func (o *Service) MakeDataBag(dataBagRequest *request.DataBagRequest) (*response.DataBagBean, error) {

	var url string
	if dataBagRequest.FileId != nil {
		url = o.dataBagURL("files", *dataBagRequest.FileId, dataBagRequest.Callback)
	} else if dataBagRequest.IntegrateId != nil {
		url = o.dataBagURL("integrations", *dataBagRequest.IntegrateId, dataBagRequest.Callback)
	} else if dataBagRequest.CompareId != nil {
		url = o.dataBagURL("comparisions", *dataBagRequest.IntegrateId, dataBagRequest.Callback)
	}
	if url == "" {
		return nil, fmt.Errorf("url is null @ DataBagService.MakeDataBag")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Put(url, headers.Header)

	var result *response.DataBagBean
	err = utils.RespToBean(resp, result)

	return result, err
}

//-------------------------------------------------------------------------------

//查询离线数据包
/***
字段		类型	必填	描述
fileId		Number	Y	通过文件转换Id创建离线数据包时必填
integrateId	Number	Y	通过集成模型Id创建离线数据包时必填
compareId	Number	Y	通过模型对比Id创建离线数据包时必填
***/
func (o *Service) GetStatus(dataBagRequest *request.DataBagRequest) ([]*response.DataBagBean, error) {

	var url string
	if dataBagRequest.FileId != nil {
		url = o.dataBagURL("files", *dataBagRequest.FileId, "")
	} else if dataBagRequest.IntegrateId != nil {
		url = o.dataBagURL("integrations", *dataBagRequest.IntegrateId, "")
	} else if dataBagRequest.CompareId != nil {
		url = o.dataBagURL("comparisions", *dataBagRequest.IntegrateId, "")
	}
	if url == "" {
		return nil, fmt.Errorf("url is null @ DataBagService.GetDataBagStatus")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(url, headers.Header)

	var result []*response.DataBagBean
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//-------------------------------------------------------------------------------

//获取离线数据包下载地址
/***
字段			类型	必填	描述
fileId			Number	Y	通过文件转换Id获取离线数据包下载地址时必填
integrateId		Number	Y	通过集成模型Id获取离线数据包下载地址时必填
compareId		Number	Y	通过模型对比Id获取离线数据包下载地址时必填
type			String	Y	值必须是“offline”
databagVersion	String	N	数据包版本，如果只有一个，则下载唯一的数据包，如果多个，则必须指定数据包版本, 例如 3.0
***/
func (o *Service) GetDownloadURL(dataBagRequest *request.DataBagRequest) (string, error) {

	var url string
	if dataBagRequest.FileId != nil {
		url = o.downloadDataBagURL("fileId", *dataBagRequest.FileId, dataBagRequest.DataBagVersion)
	} else if dataBagRequest.IntegrateId != nil {
		url = o.downloadDataBagURL("integrateId", *dataBagRequest.IntegrateId, dataBagRequest.DataBagVersion)
	} else if dataBagRequest.CompareId != nil {
		url = o.downloadDataBagURL("comapreId", *dataBagRequest.CompareId, dataBagRequest.DataBagVersion)
	}
	if url == "" {
		return "", fmt.Errorf("url is null @ DataBagService.GetDataBagDownloadURL")
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
