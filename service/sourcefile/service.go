// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

import (
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/bimface/v3/service/abstract"
)

// Service ***
type Service struct {
	*abstract.Service //base class
	supportFile       *response.FileSupportBean
}

// NewService ***
func NewService(absService *abstract.Service) *Service {
	return &Service{
		Service:     absService,
		supportFile: nil,
	}
}
