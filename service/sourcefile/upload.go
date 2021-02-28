// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

//文件上传

import (
	"fmt"

	"github.com/chnykn/bimface/bean/request"
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const (
	uploadURI              string = "/upload?name=%s"
	uploadByURLURI         string = "/upload?name=%s&url=%s"
	getFileUploadStatusURI string = "/files/%d/uploadStatus"
)

//---------------------------------------------------------------------

func (o *Service) uploadURL(fileName string, sourceId string) string {
	result := fmt.Sprintf(o.Endpoint.FileHost+uploadURI, utils.EncodeURI(fileName))
	if sourceId != "" {
		result = result + "&sourceId=" + utils.EncodeURI(sourceId)
	}
	return result
}

func (o *Service) uploadByURL(fileName, url string, sourceId string) string {
	result := fmt.Sprintf(o.Endpoint.FileHost+uploadByURLURI, utils.EncodeURI(fileName), utils.EncodeURI(url))
	if sourceId != "" {
		result = result + "&sourceId=" + utils.EncodeURI(sourceId)
	}
	return result
}

func (o *Service) getFileUploadStatusURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+getFileUploadStatusURI, fileId)
}

//------------------------------------------------------------------------------------

func (o *Service) doUploadByURL(uploadRequest *request.FileUploadRequest, token string) (*response.FileBean, error) {
	headers := utils.NewHeaders()
	headers.AddOAuth2Header(token)

	resp := o.ServiceClient.Put(o.uploadByURL(uploadRequest.Name, uploadRequest.URL,
		uploadRequest.SourceId), headers.Header)

	var result *response.FileBean
	err := utils.RespToBean(resp, result)

	return result, err
}

func (o *Service) doUploadBody(uploadRequest *request.FileUploadRequest, token string) (*response.FileBean, error) {

	/***
	data, ferr := uploadRequest.InputFile.Open()
	if ferr != nil {
		return nil, utils.NewError(ferr.Error(), "uploadRequest.InputFile.Open() @ doUploadBody")
	}
	defer data.Close()

	buf := make([]byte, uploadRequest.InputFile.Size) //uploadRequest.ContentLength
	data.Read(buf)
	***/

	//------------------------------------

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(token)
	//headers.Header["Content-Length"] = strconv.FormatInt(uploadRequest.ContentLength, 10)

	resp := o.ServiceClient.Put(o.uploadURL(uploadRequest.Name, uploadRequest.SourceId),
		headers.Header, uploadRequest.Buffer)

	var result *response.FileBean
	err := utils.RespToBean(resp, result)

	return result, err
}

//源文件相关: 上传文件
//http://static.bimface.com/book/restful/articles/api/file/upload.html
/***
字段		类型	必填	描述	示例
name		String	Y	文件的全名，使用URL编码（UTF-8），最多256个字符
sourceId	String	N	调用方的文件源Id，不能重复
url			String	N	文件的下载地址，使用URL编码（UTF-8），最多512个字符，注：在pull方式下必填，必须以http(s)://开头
***/
func (o *Service) Upload(uploadRequest *request.FileUploadRequest) (*response.FileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	if uploadRequest.IsByURL() {
		return o.doUploadByURL(uploadRequest, accessToken.Token)
		//} else if uploadRequest.IsByOSS() {
		//	return o.doUploadByOSS(uploadRequest, accessToken.Token)
	} else {
		return o.doUploadBody(uploadRequest, accessToken.Token)
	}
}

//------------------------------------------------------------------------------------

//GetFileUploadStatus
func (o *Service) GetUploadStatus(fileId int64) (*response.FileUploadStatusBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getFileUploadStatusURL(fileId), headers.Header)

	var result *response.FileUploadStatusBean
	err = utils.RespToBean(resp, result)

	return result, err
}
