// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE integrate.

package databag

import (
	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
)

func (o *Service) IntegrateDataBag(integrateId int64, callback string, keepModelDB bool) (*response.DataBagBean, error) {
	dataBagRequest := request.NewDataBagRequest()
	dataBagRequest.IntegrateId = &integrateId
	dataBagRequest.Callback = callback
	return o.MakeDataBag(dataBagRequest, keepModelDB)
}

func (o *Service) GetIntegrateStatus(integrateId int64) ([]*response.DataBagBean, error) {
	dataBagRequest := request.NewDataBagRequest()
	dataBagRequest.IntegrateId = &integrateId
	return o.GetStatus(dataBagRequest)
}

func (o *Service) GetIntegrateDownloadURL(integrateId int64) (string, error) {
	dataBagRequest := request.NewDataBagRequest()
	dataBagRequest.IntegrateId = &integrateId
	return o.GetDownloadURL(dataBagRequest)
}
