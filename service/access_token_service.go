// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
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
func (o *AccessTokenService) Get() (*response.AccessToken, error) {
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
func (o *AccessTokenService) Grant() (*response.AccessToken, error) {
	headers := utils.NewHeaders()
	headers.AddBasicAuthHeader(o.Credential.AppKey, o.Credential.AppSecret)
	resp := o.ServiceClient.Post(o.accessTokenURL(), headers.Header)

	result := response.NewAccessToken("", "")
	err := utils.RespToBean(resp, result)

	return result, err
}
