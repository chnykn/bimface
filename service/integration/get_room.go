// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	//获取单个房间信息 GET https://api.bimface.com/data/v2/integrations/{integrateId}/rooms/{roomId}
	getRoomURI string = "/data/v2/integrations/%d/rooms/%s"
)

//---------------------------------------------------------------------

func (o *Service) getRoomURL(integrateId int64, roomId string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getRoomURI, integrateId, roomId)
}

//------------------------------------------------------------------------------------

// 获取单个房间信息
func (o *Service) GetRoom(integrateId int64, roomId string) (*response.RoomBean, error) {
	result := new(response.RoomBean)
	err := o.GET(o.getRoomURL(integrateId, roomId), result)

	return result, err
}
