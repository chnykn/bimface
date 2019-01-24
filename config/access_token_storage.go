// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"time"

	"github.com/chnykn/bimface/bean/response"
)

//AccessTokenStorage ***
type AccessTokenStorage struct {
	accessToken *response.AccessToken
}

//NewAccessTokenStorage ***
func NewAccessTokenStorage() *AccessTokenStorage {
	o := &AccessTokenStorage{
		accessToken: nil,
	}
	return o
}

//Put ***
func (o *AccessTokenStorage) Put(accessToken *response.AccessToken) {
	o.accessToken = accessToken
}

//Get ***
func (o *AccessTokenStorage) Get() *response.AccessToken {

	if o.accessToken == nil {
		return nil
	}

	if maybeExpire(o.accessToken.ExpireTime) {
		return nil
	}

	return o.accessToken
}

//MaybeExpire ***
func maybeExpire(expireTime string) bool {
	expire, err := time.ParseInLocation("2006-01-02 15:04:05", expireTime, time.Local)
	if err == nil {
		return expire.Before(time.Now())
	}

	return true
}
