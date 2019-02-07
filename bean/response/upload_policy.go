// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//UploadPolicy ***
type UploadPolicy struct {
	Host         string `json:"host"`
	Policy       string `json:"policy"`
	AccessID     string `json:"accessId"`
	Signature    string `json:"signature"`
	Expire       int64  `json:"expire"`
	CallbackBody string `json:"callbackBody,omitempty"`
	ObjectKey    string `json:"objectKey"`
}

//NewUploadPolicy ***
func NewUploadPolicy() *UploadPolicy { //host string, policy string
	o := &UploadPolicy{
		// Host:   host,
		// Policy: policy,
	}
	return o
}
