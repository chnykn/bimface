// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package databag

import (
	"github.com/chnykn/bimface/v2/bean/request"
	"github.com/chnykn/bimface/v2/bean/response"
)

func (o *Service) FileDataBag(fileId int64, callback string, keepModelDB bool) (*response.DataBagBean, error) {
	dataBagRequest := request.NewDataBagRequest()
	dataBagRequest.FileId = &fileId
	dataBagRequest.Callback = callback
	return o.MakeDataBag(dataBagRequest, keepModelDB)
}

func (o *Service) GetFileStatus(fileId int64) ([]*response.DataBagBean, error) {
	dataBagRequest := request.NewDataBagRequest()
	dataBagRequest.FileId = &fileId
	return o.GetStatus(dataBagRequest)
}

func (o *Service) GetFileDownloadURL(fileId int64) (string, error) {
	dataBagRequest := request.NewDataBagRequest()
	dataBagRequest.FileId = &fileId
	return o.GetDownloadURL(dataBagRequest)
}
