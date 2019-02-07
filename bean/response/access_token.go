// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//AccessToken ***
type AccessToken struct {
	Token      string `json:"token"`
	ExpireTime string `json:"expireTime"`
}

//NewAccessToken ***
func NewAccessToken(token, expireTime string) *AccessToken {
	o := &AccessToken{
		Token:      token,
		ExpireTime: expireTime,
	}
	return o
}

// ToString get the string
func (o *AccessToken) ToString() string {
	return fmt.Sprintf("AccessToken [Token=%s, ExpireTime=%s]", o.Token, o.ExpireTime)
}
