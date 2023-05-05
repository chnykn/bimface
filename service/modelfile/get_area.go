// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	//获取单个面积分区信息 GET https://api.bimface.com/data/v2/files/{fileId}/areas/{areaId}
	getAreaURI string = "/data/v2/files/%d/areas/%s"

	//获取楼层对应面积分区列表 GET https://api.bimface.com/data/v2/files/{fileId}/areas
	getAreasURI string = "data/v2/files/%d/areas"
)

//---------------------------------------------------------------------

func (o *Service) getAreaURL(fileId int64, areaId string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getAreaURI, fileId, areaId)
}

func (o *Service) getAreasURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getAreasURI, fileId)
}

//------------------------------------------------------------------------------------

// 获取单个面积分区信息
func (o *Service) GetArea(fileId int64, areaId string) (*response.AreaBean, error) {
	result := new(response.AreaBean)
	err := o.GET(o.getAreaURL(fileId, areaId), result)

	return result, err
}

// 获取楼层对应面积分区列表
func (o *Service) GetAreas(fileId int64) ([]*response.AreaBean, error) {
	result := make([]*response.AreaBean, 0)
	err := o.GET(o.getAreasURL(fileId), &result)

	return result, err
}
