// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

//文件上传

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/bimface/v3/utils"
	"github.com/chnykn/httpkit"
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

func (o *Service) doUploadByURL(uploadRequest *request.FileUploadRequest) (*response.FileBean, error) {
	result := new(response.FileBean)
	err := o.PUT(o.uploadByURL(uploadRequest.Name, uploadRequest.URL, uploadRequest.SourceId), result)

	return result, err
}

func (o *Service) doUploadBody(uploadRequest *request.FileUploadRequest) (*response.FileBean, error) {
	result := new(response.FileBean)
	err := o.PUT(o.uploadURL(uploadRequest.Name, uploadRequest.SourceId),
		result,
		httpkit.BufferReqBody(uploadRequest.Buffer))

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

	if uploadRequest.IsByURL() {
		return o.doUploadByURL(uploadRequest)
	} else {
		return o.doUploadBody(uploadRequest)
	}
}

//------------------------------------------------------------------------------------

// GetUploadStatus **
func (o *Service) GetUploadStatus(fileId int64) (*response.FileUploadStatusBean, error) {
	result := new(response.FileUploadStatusBean)
	err := o.GET(o.getFileUploadStatusURL(fileId), result)

	return result, err
}
