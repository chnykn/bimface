// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bake

import (
	"github.com/chnykn/bimface/v3/service/abstract"
)

// Service ***
type Service struct {
	*abstract.Service //base class

}

// NewService ***
func NewService(absService *abstract.Service) *Service {
	return &Service{
		Service: absService,
	}
}
