// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

// Endpoint ***
type Endpoint struct {
	APIHost  string // consts.APIHost
	FileHost string // consts.FileHost
}

// NewEndpoint ***
func NewEndpoint(apiHost string, fileHost string) *Endpoint {
	o := &Endpoint{
		APIHost:  apiHost,
		FileHost: fileHost,
	}
	return o
}
