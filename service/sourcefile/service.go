// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

import (
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/service"
	"github.com/chnykn/bimface/utils"
)

//Service ***
type Service struct {
	service.AbstractService //base class
	AccessTokenService      *service.AccessTokenService

	supportFile *response.FileSupportBean
}

//NewService ***
func NewService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	accessTokenService *service.AccessTokenService) *Service {
	o := &Service{
		AbstractService: service.AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
		supportFile:        nil,
	}

	return o
}
