// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const supportFileURI = "/support"

var defSupportFile = &response.SupportFile{
	Length: 13,
	Types:  []string{"rvt", "rfa", "dwg", "dxf", "skp", "ifc", "dgn", "obj", "stl", "3ds", "dae", "ply", "igms"},
}

//---------------------------------------------------------------------

func (o *Service) supportFileURL() string {
	return o.Endpoint.FileHost + supportFileURI
}

//GetSupportFile ***
func (o *Service) GetSupportFile() (*response.SupportFile, error) {

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

	var result *response.SupportFile
	err = utils.RespToBean(resp, result)

	if err == nil {
		o.supportFile = result
	} else {
		o.supportFile = defSupportFile
	}

	return result, err
}
