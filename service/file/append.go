// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

//断点续传

import (
	"fmt"
	"mime/multipart"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const (
	//创建追加文件
	createAppendFileURI string = "/appendFiles?name=%s&length=%d" //&sourceId=%s
	getAppendFileURI    string = "/appendFiles/%d"
	uploadAppendFileURI string = "/appendFiless/%d/data?position=%d"
)

//---------------------------------------------------------------------

func (o *Service) createAppendFileURL(fileName string, length int64, sourceId string) string {
	result := fmt.Sprintf(o.Endpoint.FileHost+createAppendFileURI, fileName, length) //&sourceId=%s
	if sourceId != "" {
		result = result + "&sourceId=" + sourceId
	}
	return result
}

func (o *Service) getAppendFileURL(appendFileId int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+getAppendFileURI, appendFileId)
}

func (o *Service) uploadAppendFileURL(appendFileId int64, position int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+uploadAppendFileURI, appendFileId, position)
}

//---------------------------------------------------------------------

//断点续传: 创建追加文件
/***
字段		类型	必填	描述
name		String	Y	文件的全名，使用URL编码（UTF-8），最多256个字符
sourceId	String	N	调用方的文件源Id，不能重复
length		Number	Y	上传文件长度
***/
func (o *Service) createAppendFile(fileName string, length int64, sourceId string) (*response.AppendFileBean, error) {

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
	supportFile, err = o.GetSupportFile()
	if err != nil {
		return nil, err
	}

	err = utils.CheckFileType(supportFile.Types, fileName)
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Post(o.createAppendFileURL(fileName, length, sourceId), headers.Header)

	var result *response.AppendFileBean
	err = utils.RespToBean(resp, result)

	return result, err
}

//断点续传: 查询追加文件信息
/***
字段			类型	必填	描述
appendFileId	Number	Y	append file id
***/
func (o *Service) GetAppendFileWithAccessToken(appendFileId int64, token string) (*response.AppendFileBean, error) {
	headers := utils.NewHeaders()
	headers.AddOAuth2Header(token)

	resp := o.ServiceClient.Get(o.getAppendFileURL(appendFileId), headers.Header)

	var result *response.AppendFileBean
	err := utils.RespToBean(resp, result)

	return result, err
}

//GetAppendFile same to GetAppendFileWithAccessToken
/***
字段			类型	必填	描述
appendFileId	Number	Y	append file id
***/
func (o *Service) GetAppendFile(appendFileId int64) (*response.AppendFileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	return o.GetAppendFileWithAccessToken(appendFileId, accessToken.Token)
}

//断点续传: 追加上传
/***
字段			类型	必填	描述
appendFileId	Number	Y	追加文件id
position		Number	N	追加上传开始位置，默认为0
***/
func (o *Service) UploadAppendFile(file *multipart.FileHeader, appendFileId int64) (*response.AppendFileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	var appendFile *response.AppendFileBean
	appendFile, err = o.GetAppendFileWithAccessToken(appendFileId, accessToken.Token)
	if err != nil {
		return nil, err
	}

	//------------------------------

	data, ferr := file.Open()
	if ferr != nil {
		return nil, ferr
	}
	defer data.Close()

	len1 := file.Size - appendFile.Position
	buf := make([]byte, len1)
	_, _ = data.ReadAt(buf, appendFile.Position)

	//------------------------------

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Post(o.uploadAppendFileURL(appendFileId, appendFile.Position), headers.Header, buf)

	var result *response.AppendFileBean
	err = utils.RespToBean(resp, result)

	return result, err
}
