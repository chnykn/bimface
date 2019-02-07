// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

const supportFileURI = "/support"

var defSupportFile = &response.SupportFile{
	Length: 13,
	Types:  []string{"rvt", "rfa", "dwg", "dxf", "skp", "ifc", "dgn", "obj", "stl", "3ds", "dae", "ply", "igms"},
}

//SupportFileService ***
type SupportFileService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewSupportFileService ***
func NewSupportFileService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *SupportFileService {
	o := &SupportFileService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *SupportFileService) supportFileURL() string {
	return o.Endpoint.FileHost + supportFileURI
}

//GetSupportWithAccessToken ***
func (o *SupportFileService) GetSupportWithAccessToken(token string) (*response.SupportFile, error) {

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(token)

	resp := o.ServiceClient.Get(o.supportFileURL(), headers.Header)

	result := response.NewSupportFile()
	err := utils.RespToBean(resp, result)

	return result, err

	//return defSupportFile, nil
}

//GetSupport ***
func (o *SupportFileService) GetSupport() (*response.SupportFile, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	return o.GetSupportWithAccessToken(accessToken.Token)

	//return defSupportFile, nil
}
