// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

import (
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const supportFileURI = "/support"

var defSupportFile = &response.FileSupportBean{
	Length: 13,
	Types:  []string{"rvt", "rfa", "dwg", "dxf", "skp", "ifc", "dgn", "obj", "stl", "3ds", "dae", "ply", "igms"},
}

//---------------------------------------------------------------------

func (o *Service) supportFileURL() string {
	return o.Endpoint.FileHost + supportFileURI
}

//GetSupportFile ***
func (o *Service) GetSupportFile() (*response.FileSupportBean, error) {

	if o.supportFile != nil {
		return o.supportFile, nil
	}

	//---------------------

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.supportFileURL(), headers.Header)

	result := new(response.FileSupportBean)
	err = utils.RespToBean(resp, result)

	if err == nil {
		o.supportFile = result
	} else {
		o.supportFile = defSupportFile
	}

	return result, err
}
