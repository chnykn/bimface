// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bake

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/bimface/v3/utils"
	"github.com/chnykn/httpkit"
)

const (
	bakeURI string = "/%s/%d/bake" //%s = files or integrations; %d = fileId or integrateId

)

func (o *Service) bakeURL(kind string, objectId int64, callback string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+bakeURI, kind, objectId)

	if callback != "" {
		result = result + "?callback=" + utils.EncodeURI(callback)
	}

	return result
}

// ---------------------------------------------------------------------
// kind must in [files, integrations]
func (o *Service) Bake(kind string, objectId int64, config map[string]string, callback string) (*response.DataBagBean, error) {

	result := new(response.DataBagBean)

	body := make(map[string]interface{})
	body["config"] = config

	err := o.PUT(o.bakeURL(kind, objectId, callback), result, httpkit.JsonReqBody(body))
	return result, err
}

func (o *Service) GetBakeStatus(kind string, objectId int64) ([]*response.DataBagBean, error) {
	result := make([]*response.DataBagBean, 0)
	err := o.GET(o.bakeURL(kind, objectId, ""), result)

	return result, err
}
