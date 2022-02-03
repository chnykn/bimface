// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

//获取文件信息

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	fileInfoURI string = "/files/%d"
	fileListURI string = "/files"
)

//---------------------------------------------------------------------

func (o *Service) fileMetaURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+fileInfoURI, fileId)
}

//offset   : 查询结果偏移，从查询结果的第offset条开始返回数据
//rows     : 查询结果数, 默认为100， 最大500
//status   : 文件状态,uploading，success，failure
//startTime: 起始日期,格式为 yyyy-MM-dd
//endTime  : 截止日期,格式为 yyyy-MM-dd
func (o *Service) fileListURL(offset, rows int, status string, startTime, endTime string) string {

	params := ""
	if offset > 0 {
		params = params + "offset=" + strconv.Itoa(offset)
	}
	if rows > 0 {
		params = params + "&rows=" + strconv.Itoa(rows)
	}
	if status != "" {
		params = params + "&status=" + status
	}
	if startTime != "" {
		params = params + "&startTime=" + startTime
	}
	if endTime != "" {
		params = params + "&endTime=" + endTime
	}

	result := fmt.Sprintf(o.Endpoint.FileHost + fileListURI)
	params = strings.TrimLeft(params, "&")

	return result + "?" + params
}

//获取文件信息
func (o *Service) GetInfo(fileId int64) (*response.FileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.fileMetaURL(fileId), headers.Header)

	result := new(response.FileBean)
	err = utils.RespToBean(resp, result)

	return result, err
}

// 获取文件信息列表 GET https://file.bimface.com/files
// http://static.bimface.com/restful-apidoc/dist/file.html#_listfilesusingget
func (o *Service) GetList() ([]*response.FileBean, error) {
	return o.GetListEx(0, 0, "", "", "")
}

// 获取文件信息列表 GET https://file.bimface.com/files
//offset   : 查询结果偏移，从查询结果的第offset条开始返回数据
//rows     : 查询结果数, 默认为100， 最大500
//status   : 文件状态,uploading，success，failure
//startTime: 起始日期,格式为 yyyy-MM-dd
//endTime  : 截止日期,格式为 yyyy-MM-dd
func (o *Service) GetListEx(offset, rows int, status string, startTime, endTime string) ([]*response.FileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(
		o.fileListURL(offset, rows, status, startTime, endTime),
		headers.Header)

	result := make([]*response.FileBean, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}
