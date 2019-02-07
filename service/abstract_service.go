// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

//AbstractService ***
type AbstractService struct {
	Endpoint      *config.Endpoint
	ServiceClient *utils.ServiceClient
}
