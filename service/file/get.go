// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

//获取文件信息

import (
	"fmt"

	"github.com/chnykn/bimface/bean/response"

	"github.com/chnykn/bimface/utils"
)

const (
	getFileMetaURI string = "/files/%d"
	getFileListURI string = "/files"
)

//---------------------------------------------------------------------

func (o *Service) getFileMetaURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+getFileMetaURI, fileId)
}

func (o *Service) getFileListURL() string {
	return fmt.Sprintf(o.Endpoint.FileHost + getFileListURI)
}

//GetFileMeta 源文件相关: 获取文件信息
func (o *Service) GetFileMeta(fileId int64) (*response.FileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getFileMetaURL(fileId), headers.Header)

	result := response.NewFileBean()
	err = utils.RespToBean(resp, result)

	return result, err
}

// GetFileList 获取文件信息列表 GET https://file.bimface.com/files
// http://static.bimface.com/restful-apidoc/dist/file.html#_listfilesusingget
func (o *Service) GetFileList() ([]*response.FileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getFileListURL(), headers.Header)

	var result []*response.FileBean //:= make([]*response.FileBean, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}
