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
	//获取指定面积分区的属性 GET https://api.bimface.com/data/v2/integrations/{integrateId}/areas/{areaId}
	getAreaURI string = "/data/v2/integrations/%d/areas/%s"

	//获取楼层对应面积分区列表 GET https://api.bimface.com/data/v2/integrations/{integrateId}/areas
	getAreasURI string = "data/v2/integrations/%d/areas"
)

//---------------------------------------------------------------------

func (o *Service) getAreaURL(integrateId int64, areaId string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getAreaURI, integrateId, areaId)
}

func (o *Service) getAreasURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getAreasURI, integrateId)
}

//------------------------------------------------------------------------------------

//获取指定面积分区的属性
func (o *Service) GetArea(integrateId int64, areaId string) (*response.AreaBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getAreaURL(integrateId, areaId), headers.Header)

	var result *response.AreaBean
	err = utils.RespToBean(resp, result)

	return result, err
}

//获取楼层对应面积分区列表
func (o *Service) GetAreas(integrateId int64) ([]*response.AreaBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getAreasURL(integrateId), headers.Header)

	var result []*response.AreaBean
	err = utils.RespToBean(resp, &result)

	return result, err
}
