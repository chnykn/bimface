// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	//为文件创建bake数据包 PUT https://api.bimface.com/files/{fileId}/bake
	bakeURI = "files/%d/bake"
)

func (o *Service) bakeURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+bakeURI, fileId)
}

//------------------------------------------------------------------

func (o *Service) Bake(fileId int64, config map[string]string) (*response.DataBagBean, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := make(map[string]interface{})
	body["config"] = config

	resp := o.ServiceClient.Put(o.bakeURL(fileId), headers.Header, req.BodyJSON(body))

	var result *response.DataBagBean
	err = utils.RespToBean(resp, result)

	return result, err
}

func (o *Service) DefaultBake(fileId int64) (*response.DataBagBean, error) {

	body := make(map[string]string)
	body["texture"] = "true"

	return o.Bake(fileId, body)
}

//---------------------------------------

//查询bake数据包
func (o *Service) GetBakeDataBag(fileId int64) ([]*response.DataBagBean, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.bakeURL(fileId), headers.Header)

	var result []*response.DataBagBean
	err = utils.RespToBean(resp, &result)

	return result, err
}
