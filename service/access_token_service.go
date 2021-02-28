// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/config"
	"github.com/chnykn/bimface/v2/utils"
)

//获取AccessToken
const accessTokenURI string = "/oauth2/token"

//AccessTokenService ***
type AccessTokenService struct {
	AbstractService //base class

	Credential         *config.Credential
	AccessTokenStorage *config.AccessTokenStorage
}

//NewAccessTokenService 获取AccessToken
//http://static.bimface.com/book/restful/articles/api/accesstoken.html
func NewAccessTokenService(serviceClient *utils.ServiceClient,
	endpoint *config.Endpoint, credential *config.Credential) *AccessTokenService {
	o := &AccessTokenService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		Credential:         credential,
		AccessTokenStorage: config.NewAccessTokenStorage(),
	}

	return o
}

//---------------------------------------------------------------------

func (o *AccessTokenService) accessTokenURL() string {
	return o.Endpoint.APIHost + accessTokenURI
}

//Get ***
func (o *AccessTokenService) Get() (*response.AccessTokenBean, error) {
	accessToken := o.AccessTokenStorage.Get()
	var err error

	if accessToken == nil {
		accessToken, err = o.Grant()
		if err == nil {
			o.AccessTokenStorage.Put(accessToken)
		}
	}
	return accessToken, err
}

//Grant ***
func (o *AccessTokenService) Grant() (*response.AccessTokenBean, error) {
	headers := utils.NewHeaders()
	headers.AddBasicAuthHeader(o.Credential.AppKey, o.Credential.AppSecret)
	resp := o.ServiceClient.Post(o.accessTokenURL(), headers.Header)

	var result *response.AccessTokenBean
	err := utils.RespToBean(resp, result)

	return result, err
}
