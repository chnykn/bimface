// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/httpkit"
)

const (
	translateURI        string = "/translate"
	getTranslateURI     string = "/translate?fileId=%d"
	translateDetailsURI string = "/translateDetails"
)

//---------------------------------------------------------------------

func (o *Service) translateURL() string {
	return o.Endpoint.APIHost + translateURI
}

func (o *Service) getTranslateURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getTranslateURI, fileId)
}

func (o *Service) translateDetailsURL() string {
	return fmt.Sprintf(o.Endpoint.APIHost + translateDetailsURI)
}

//-----------------------------------------------------------------------------------

// 发起文件转换
func (o *Service) Translate(transRequest *request.FileTranslateRequest) (*response.FileTranslateBean, error) {
	result := new(response.FileTranslateBean)

	body := httpkit.JsonReqBody(transRequest)
	err := o.PUT(o.translateURL(), result, body)

	return result, err
}

//-----------------------------------------------------------------------------------

// 获取转换状态
func (o *Service) GetStatus(fileId int64) (*response.FileTranslateBean, error) {
	result := new(response.FileTranslateBean)
	err := o.GET(o.getTranslateURL(fileId), result)

	return result, err
}

//-----------------------------------------------------------------------------------

// 批量获取转换状态详情
func (o *Service) GetDetails(queryRequest *request.TranslateQueryRequest) (*response.FileTranslateDetailBeanPageList, error) {

	var err error
	result := new(response.FileTranslateDetailBeanPageList)

	if queryRequest != nil {
		body := httpkit.JsonReqBody(queryRequest)
		err = o.POST(o.translateDetailsURL(), result, body)
	} else {
		err = o.POST(o.translateDetailsURL(), result)
	}

	return result, err
}
