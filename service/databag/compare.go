// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE compare.

package databag

import (
	"github.com/chnykn/bimface/bean/request"
	"github.com/chnykn/bimface/bean/response"
)

func (o *Service) CompareDataBag(compareId int64, callback string) (*response.DataBagBean, error) {
	dataBagRequest := request.NewDataBagRequest()
	dataBagRequest.CompareId = &compareId
	dataBagRequest.Callback = callback
	return o.MakeDataBag(dataBagRequest)
}

func (o *Service) GetCompareStatus(compareId int64) ([]*response.DataBagBean, error) {
	dataBagRequest := request.NewDataBagRequest()
	dataBagRequest.CompareId = &compareId
	return o.GetStatus(dataBagRequest)
}

func (o *Service) GetCompareDownloadURL(compareId int64) (string, error) {
	dataBagRequest := request.NewDataBagRequest()
	dataBagRequest.CompareId = &compareId
	return o.GetDownloadURL(dataBagRequest)
}
