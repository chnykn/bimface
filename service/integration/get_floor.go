// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const (
	//获取楼层列表 GET https://api.bimface.com/data/v2/integrations/{integrateId}/floors
	floorsURI string = "/data/v2/integrations/%d/floors"
)

//---------------------------------------------------------------------

func (o *Service) floorsURL(integrateId int64, includeArea, includeRoom bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+floorsURI, integrateId)

	var s = "?"
	if includeArea {
		result = result + "?includeArea=true"
		s = "&"
	}

	if includeRoom {
		result = result + s + "includeRoom=true"
	}

	return result
}

//---------------------------------------------------------------------

//获取单模型的楼层信息
func (o *Service) GetFloors(integrateId int64, includeArea, includeRoom bool) ([]*response.FloorBean, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.floorsURL(integrateId, includeArea, includeRoom), headers.Header)

	var result []*response.FloorBean
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
