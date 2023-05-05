// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bake

import (
	"github.com/chnykn/bimface/v3/bean/response"
)

func (o *Service) FileBake(fileId int64, config map[string]string, callback string) (*response.DataBagBean, error) {
	return o.Bake("files", fileId, config, callback)
}
func (o *Service) FileDefaultBake(fileId int64, callback string) (*response.DataBagBean, error) {
	config := make(map[string]string)
	config["texture"] = "true"

	return o.Bake("files", fileId, config, callback)
}

func (o *Service) GetFileStatus(fileId int64) ([]*response.DataBagBean, error) {
	return o.GetBakeStatus("files", fileId)
}
