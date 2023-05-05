// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	//获取MEP系统信息 GET https://api.bimface.com/data/v2/files/{fileId}/MEPSystem
	mepSystemURI string = "data/v2/files/%d/MEPSystem"
)

//---------------------------------------------------------------------

func (o *Service) mepSystemURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+mepSystemURI, fileId)
}

//---------------------------------------------------------------------

// 获取MEP系统信息
func (o *Service) GetMEPSystem(fileId int64) (*response.MEPSysBean, error) {
	result := new(response.MEPSysBean)
	err := o.GET(o.mepSystemURL(fileId), result)

	return result, err
}
