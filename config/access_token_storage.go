// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"time"

	"github.com/chnykn/bimface/v2/bean/response"
)

type AccessTokenStorage struct {
	accessToken *response.AccessTokenBean
}

func NewAccessTokenStorage() *AccessTokenStorage {
	o := &AccessTokenStorage{
		accessToken: nil,
	}
	return o
}

func (o *AccessTokenStorage) Put(accessToken *response.AccessTokenBean) {
	o.accessToken = accessToken
}

func (o *AccessTokenStorage) Get() *response.AccessTokenBean {

	if o.accessToken == nil {
		return nil
	}

	if maybeExpire(o.accessToken.ExpireTime) {
		return nil
	}

	return o.accessToken
}

func maybeExpire(expireTime string) bool {
	expire, err := time.ParseInLocation("2006-01-02 15:04:05", expireTime, time.Local)
	if err == nil {
		return expire.Before(time.Now())
	}

	return true
}
