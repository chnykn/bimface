// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/utils"
)

//获取文件下载地址
const downloadURI string = "/download/url?fileId=%d" //&fileName=%s  fileName? name?

//---------------------------------------------------------------------

func (o *Service) downloadURL(fileId int64, fileName string) string {
	result := fmt.Sprintf(o.Endpoint.FileHost+downloadURI, fileId)
	if fileName != "" {
		result = result + "&name=" + fileName
	}
	return result
}

//---------------------------------------------------------------------

//GetDownloadURL 源文件相关: 获取文件下载地址
/***
字段	类型	必填	描述
fileId	Number	Y	文件Id
name	String	N	自定义文件下载名
***/
func (o *Service) GetDownloadURL(fileId int64, fileName string) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.downloadURL(fileId, fileName), headers.Header)

	result := new(string)
	err = utils.RespToBean(resp, result)

	return *result, err
}

//GetDownloadResp 下载文件
/***
字段	类型	必填	描述
fileId	Number	Y	文件Id
name	String	N	自定义文件下载名
***/
func (o *Service) GetDownloadResp(fileId int64, fileName string) (*req.Resp, error) {
	fileURL, err := o.GetDownloadURL(fileId, fileName)

	if err == nil {
		resp := o.ServiceClient.Get(fileURL)
		return resp, nil
	}

	return nil, err
}
