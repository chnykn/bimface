// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
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

// 获取指定面积分区的属性
func (o *Service) GetArea(integrateId int64, areaId string) (*response.AreaBean, error) {
	result := new(response.AreaBean)
	err := o.GET(o.getAreaURL(integrateId, areaId), result)

	return result, err
}

// 获取楼层对应面积分区列表
func (o *Service) GetAreas(integrateId int64) ([]*response.AreaBean, error) {
	result := make([]*response.AreaBean, 0)
	err := o.GET(o.getAreasURL(integrateId), &result)

	return result, err
}
