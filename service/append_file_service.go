// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"mime/multipart"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

const (
	createAppendFileURI string = "/appendFiles?name=%s&length=%d" //&sourceId=%s
	queryAppendFileURI  string = "/appendFiles/%d"
	uploadAppendFileURI string = "/appendFiless/%d/data?position=%d"
)

//AppendFileService ***
type AppendFileService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
	SupportFileService *SupportFileService
}

//NewAppendFileService ***
func NewAppendFileService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService,
	supportFileService *SupportFileService) *AppendFileService {
	o := &AppendFileService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
		SupportFileService: supportFileService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *AppendFileService) createAppendFileURL(fileName string, length int64, sourceID string) string {
	result := fmt.Sprintf(o.Endpoint.FileHost+createAppendFileURI, fileName, length) //&sourceId=%s
	if sourceID != "" {
		result = result + "&sourceId=" + sourceID
	}
	return result
}

func (o *AppendFileService) queryAppendFileURL(appendFileID int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+queryAppendFileURI, appendFileID)
}

func (o *AppendFileService) uploadAppendFileURL(appendFileID int64, position int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+uploadAppendFileURI, appendFileID, position)
}

//---------------------------------------------------------------------

//GetSupport 断点续传: 创建追加文件
//http://static.bimface.com/book/restful/articles/api/append/create-appendfile.html
/***
字段		类型	必填	描述
name		String	Y	文件的全名，使用URL编码（UTF-8），最多256个字符
sourceId	String	N	调用方的文件源ID，不能重复
length		Number	Y	上传文件长度
***/
func (o *AppendFileService) createAppendFile(fileName string, length int64, sourceID string) (*response.AppendFile, error) {

	err := utils.CheckFileName(fileName)
	if err != nil {
		return nil, err
	}

	var accessToken *response.AccessToken
	accessToken, err = o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	var supportFile *response.SupportFile
	supportFile, err = o.SupportFileService.GetSupportWithAccessToken(accessToken.Token)
	if err != nil {
		return nil, err
	}

	err = utils.CheckFileType(supportFile.Types, fileName)
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Post(o.createAppendFileURL(fileName, length, sourceID), headers.Header)

	result := response.NewAppendFile()
	err = utils.RespToBean(resp, result)

	return result, err
}

//QueryAppendFileWithAccessToken 断点续传: 查询追加文件信息
//http://static.bimface.com/book/restful/articles/api/append/query-appendfile.html
/***
字段			类型	必填	描述
appendFileId	Number	Y	append file id
***/
func (o *AppendFileService) QueryAppendFileWithAccessToken(appendFileID int64, token string) (*response.AppendFile, error) {
	headers := utils.NewHeaders()
	headers.AddOAuth2Header(token)

	resp := o.ServiceClient.Get(o.queryAppendFileURL(appendFileID), headers.Header)

	result := response.NewAppendFile()
	err := utils.RespToBean(resp, result)

	return result, err
}

//QueryAppendFile same to QueryAppendFileWithAccessToken
/***
字段			类型	必填	描述
appendFileId	Number	Y	append file id
***/
func (o *AppendFileService) QueryAppendFile(appendFileID int64) (*response.AppendFile, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	return o.QueryAppendFileWithAccessToken(appendFileID, accessToken.Token)
}

//UploadAppendFile 断点续传: 追加上传
//http://static.bimface.com/book/restful/articles/api/append/upload-appendfile.html
/***
字段			类型	必填	描述
appendFileId	Number	Y	追加文件id
position		Number	N	追加上传开始位置，默认为0
***/
func (o *AppendFileService) UploadAppendFile(file *multipart.FileHeader, appendFileID int64) (*response.AppendFile, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	var appendFile *response.AppendFile
	appendFile, err = o.QueryAppendFileWithAccessToken(appendFileID, accessToken.Token)
	if err != nil {
		return nil, err
	}

	//------------------------------

	data, ferr := file.Open()
	if ferr != nil {
		return nil, ferr
	}
	defer data.Close()

	len := file.Size - appendFile.Position
	buf := make([]byte, len)
	data.ReadAt(buf, appendFile.Position)

	//------------------------------

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Post(o.uploadAppendFileURL(appendFileID, appendFile.Position), headers.Header, buf)

	result := response.NewAppendFile()
	err = utils.RespToBean(resp, result)

	return result, err
}
