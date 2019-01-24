// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

// Credential ***
type Credential struct {
	AppKey    string
	AppSecret string
}

//NewCredential ***
func NewCredential(appKey string, appSecret string) *Credential {
	o := &Credential{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	return o
}
