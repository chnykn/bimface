// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bake

import (
	"fmt"

	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
	"github.com/imroc/req"
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

//---------------------------------------------------------------------
// kind must in [files, integrations]
func (o *Service) Bake(kind string, objectId int64, config map[string]string, callback string) (*response.DataBagBean, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := make(map[string]interface{})
	body["config"] = config

	resp := o.ServiceClient.Put(o.bakeURL(kind, objectId, callback), headers.Header, req.BodyJSON(body))

	result := new(response.DataBagBean)
	err = utils.RespToBean(resp, result)

	return result, err
}

func (o *Service) GetBakeStatus(kind string, objectId int64) ([]*response.DataBagBean, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.bakeURL(kind, objectId, ""), headers.Header)

	result := make([]*response.DataBagBean, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}
