// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abstract

import (
	"github.com/chnykn/bimface/v3/config"
	"github.com/chnykn/httpkit"
)

// Service ***
type Service struct {
	Endpoint   *config.Endpoint
	Credential *config.Credential

	authToken *authToken
	httpClint *httpkit.Client
}

func NewService(credential *config.Credential, endpoint *config.Endpoint) *Service {
	res := &Service{
		Endpoint:   endpoint,
		Credential: credential,
		authToken:  nil,
	}
	res.httpClint = httpkit.NewClient(res.beforeReq)

	return res
}
