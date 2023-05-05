// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
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

// 获取单模型的楼层信息
func (o *Service) GetFloors(integrateId int64, includeArea, includeRoom bool) ([]*response.FloorBean, error) {
	result := make([]*response.FloorBean, 0)
	err := o.GET(o.floorsURL(integrateId, includeArea, includeRoom), result)

	return result, err
}
