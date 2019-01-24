// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//UploadPolicy ***
type UploadPolicy struct {
	Host         string
	Policy       string
	AccessID     string
	Signature    string
	Expire       int64
	CallbackBody string
	ObjectKey    string
}

//NewUploadPolicy ***
func NewUploadPolicy() *UploadPolicy { //host string, policy string
	o := &UploadPolicy{
		// Host:   host,
		// Policy: policy,
	}
	return o
}
