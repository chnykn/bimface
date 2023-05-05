// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/httpkit"
)

const (
	//发起模型集成 PUT    https://api.bimface.com/integrate
	integrateURI string = "/integrate"

	//查询集成状态 GET    https://api.bimface.com/integrate
	//删除集成模型 DELETE https://api.bimface.com/integrate
	integrateWithIdURI string = "/integrate?integrateId=%d"

	//批量查询集成模型状态详情 POST https://api.bimface.com/integrateDetails
	integrateDetailsURI string = "/integrateDetails"
)

//---------------------------------------------------------------------

func (o *Service) integrateURL() string {
	return o.Endpoint.APIHost + integrateURI
}

func (o *Service) integrateWithIdURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+integrateWithIdURI, integrateId)
}

func (o *Service) integrateDetailsURL() string {
	return fmt.Sprintf(o.Endpoint.APIHost + integrateDetailsURI)
}

//---------------------------------------------------------------------

func (o *Service) Integrate(integrateRequest *request.FileIntegrateRequest) (*response.FileIntegrateBean, error) {
	result := new(response.FileIntegrateBean)

	body := httpkit.JsonReqBody(integrateRequest)
	err := o.PUT(o.integrateURL(), result, body)

	return result, err
}

// 获取转换状态
func (o *Service) GetStatus(integrateId int64) (*response.FileIntegrateBean, error) {
	result := new(response.FileIntegrateBean)
	err := o.GET(o.integrateWithIdURL(integrateId), result)

	return result, err
}

// 获取转换状态
func (o *Service) Delete(integrateId int64) error {
	err := o.DELETE(o.integrateWithIdURL(integrateId), nil)

	return err
}

//---------------------------------------------------------------------

// 批量查询集成模型状态详情
func (o *Service) GetDetails(queryRequest *request.IntegrateQueryRequest) (*response.FileIntegrateDetailBeanPageList, error) {

	var err error
	var result = new(response.FileIntegrateDetailBeanPageList)

	if queryRequest != nil {
		body := httpkit.JsonReqBody(queryRequest)
		err = o.POST(o.integrateDetailsURL(), result, body)
	} else {
		err = o.POST(o.integrateDetailsURL(), result)
	}

	return result, err
}
