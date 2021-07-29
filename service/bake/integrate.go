// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE integrate.

package bake

import (
	"github.com/chnykn/bimface/v2/bean/response"
)

func (o *Service) IntegrateBake(integrateId int64, config map[string]string, callback string) (*response.DataBagBean, error) {
	return o.Bake("integrations", integrateId, config, callback)
}

func (o *Service) IntegrateDefaultBake(integrateId int64, callback string) (*response.DataBagBean, error) {
	config := make(map[string]string)
	config["texture"] = "true"

	return o.Bake("integrations", integrateId, config, callback)
}

func (o *Service) GetIntegrateStatus(integrateId int64) ([]*response.DataBagBean, error) {
	return o.GetBakeStatus("integrations", integrateId)
}
