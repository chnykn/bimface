// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"

	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	//获取MEP系统信息 GET https://api.bimface.com/data/v2/files/{fileId}/MEPSystem
	mepSystemURI string = "data/v2/files/%d/MEPSystem"
)

//---------------------------------------------------------------------

func (o *Service) mepSystemURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+mepSystemURI, fileId)
}

//---------------------------------------------------------------------

//获取MEP系统信息
func (o *Service) GetMEPSystem(fileId int64) (*response.MEPSysBean, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.mepSystemURL(fileId), headers.Header)

	result := new(response.MEPSysBean)
	err = utils.RespToBean(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
