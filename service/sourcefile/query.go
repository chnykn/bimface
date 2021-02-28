// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

//获取文件信息

import (
	"fmt"

	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	fileMetaURI string = "/files/%d"
	fileListURI string = "/files"
)

//---------------------------------------------------------------------

func (o *Service) fileMetaURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+fileMetaURI, fileId)
}

func (o *Service) fileListURL() string {
	return fmt.Sprintf(o.Endpoint.FileHost + fileListURI)
}

//获取文件信息
func (o *Service) GetMeta(fileId int64) (*response.FileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.fileMetaURL(fileId), headers.Header)

	var result *response.FileBean
	err = utils.RespToBean(resp, result)

	return result, err
}

// 获取文件信息列表 GET https://file.bimface.com/files
// http://static.bimface.com/restful-apidoc/dist/file.html#_listfilesusingget
func (o *Service) GetList() ([]*response.FileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.fileListURL(), headers.Header)

	var result []*response.FileBean //:= make([]*response.FileBean, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}
