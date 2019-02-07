// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"strconv"

	"github.com/chnykn/bimface/bean/request"
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

const (
	uploadURI      string = "/upload?name=%s"
	uploadByURLURI string = "/upload?name=%s&url=%s"
	uploadByOssURI string = "/upload?name=%s&bucket=%s&objectKey=%s"

	//getUploadPolicyURI string = "/upload/policy?name=%s"

	deleteFileURI      string = "/file?fileId=%d"
	getFileMetadataURI string = "/metadata?fileId=%d"
)

//UploadService ***
type UploadService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewUploadService ***
func NewUploadService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *UploadService {
	o := &UploadService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *UploadService) uploadURL(fileName string, sourceID string) string {
	result := fmt.Sprintf(o.Endpoint.FileHost+uploadURI, fileName)
	if sourceID != "" {
		result = result + "&sourceId=" + utils.EncodeURI(sourceID)
	}
	return result
}

func (o *UploadService) uploadByURL(fileName, url string, sourceID string) string {
	result := fmt.Sprintf(o.Endpoint.FileHost+uploadByURLURI, fileName, url)
	if sourceID != "" {
		result = result + "&sourceId=" + utils.EncodeURI(sourceID)
	}
	return result
}

func (o *UploadService) uploadByOssURL(fileName, bucket, objectKey string, sourceID string) string {
	result := fmt.Sprintf(o.Endpoint.FileHost+uploadByOssURI, fileName, bucket, objectKey)
	if sourceID != "" {
		result = result + "&sourceId=" + utils.EncodeURI(sourceID)
	}
	return result
}

/**
func (o *UploadService) getUploadPolicyURL(fileName string) string {
	return fmt.Sprintf(o.Endpoint.FileHost+getUploadPolicyURI, fileName)
}
**/

func (o *UploadService) deleteFileURL(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+deleteFileURI, fileID)
}

func (o *UploadService) getFileMetadataURL(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+getFileMetadataURI, fileID)
}

//------------------------------------------------------------------------------------

func (o *UploadService) doUploadByURL(uploadRequest *request.UploadRequest, token string) (*response.FileBean, error) {
	headers := utils.NewHeaders()
	headers.AddOAuth2Header(token)

	resp := o.ServiceClient.Put(o.uploadByURL(uploadRequest.Name, uploadRequest.URL,
		uploadRequest.SourceID), headers.Header)

	result := response.NewFileBean()
	err := utils.RespToBean(resp, result)

	return result, err
}

func (o *UploadService) doUploadByOSS(uploadRequest *request.UploadRequest, token string) (*response.FileBean, error) {
	headers := utils.NewHeaders()
	headers.AddOAuth2Header(token)

	resp := o.ServiceClient.Put(o.uploadByOssURL(uploadRequest.Name, uploadRequest.Bucket,
		uploadRequest.ObjectKey, uploadRequest.SourceID), headers.Header)

	result := response.NewFileBean()
	err := utils.RespToBean(resp, result)

	return result, err
}

func (o *UploadService) doUploadBody(uploadRequest *request.UploadRequest, token string) (*response.FileBean, error) {

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
	headers.Header["Content-Length"] = strconv.FormatInt(uploadRequest.ContentLength, 10)

	resp := o.ServiceClient.Put(o.uploadURL(uploadRequest.Name, uploadRequest.SourceID),
		headers.Header, uploadRequest.InputStream)

	result := response.NewFileBean()
	err := utils.RespToBean(resp, result)

	return result, err
}

//Upload 源文件相关: 上传文件
//http://static.bimface.com/book/restful/articles/api/file/upload.html
/***
字段		类型	必填	描述	示例
name		String	Y	文件的全名，使用URL编码（UTF-8），最多256个字符
sourceId	String	N	调用方的文件源ID，不能重复
url			String	N	文件的下载地址，使用URL编码（UTF-8），最多512个字符，注：在pull方式下必填，必须以http(s)://开头
***/
func (o *UploadService) Upload(uploadRequest *request.UploadRequest) (*response.FileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	if uploadRequest.IsByURL() {
		return o.doUploadByURL(uploadRequest, accessToken.Token)
	} else if uploadRequest.IsByOSS() {
		return o.doUploadByOSS(uploadRequest, accessToken.Token)
	} else {
		return o.doUploadBody(uploadRequest, accessToken.Token)
	}
}

//------------------------------------------------------------------------------------

//DeleteFile 源文件相关: 删除文件
//http://static.bimface.com/book/restful/articles/api/file/delete-file.html
func (o *UploadService) DeleteFile(fileID int64) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Delete(o.deleteFileURL(fileID), headers.Header)

	result, err := utils.RespToResult(resp)
	if err != nil {
		return result.Code, err
	}

	return result.Code, nil
}

//GetFileMetadata 源文件相关: 获取文件元信息
//http://static.bimface.com/book/restful/articles/api/file/get-file-metadata.html
func (o *UploadService) GetFileMetadata(fileID int64) (*response.FileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getFileMetadataURL(fileID), headers.Header)

	result := response.NewFileBean()
	err = utils.RespToBean(resp, result)

	return result, err
}
