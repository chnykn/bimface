// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chnykn/bimface/v2/config"
	"github.com/chnykn/bimface/v2/utils"
)

//AbstractService ***
type AbstractService struct {
	Endpoint      *config.Endpoint
	ServiceClient *utils.ServiceClient
}
