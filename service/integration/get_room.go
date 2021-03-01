// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
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

//获取单个房间信息
func (o *Service) GetRoom(integrateId int64, roomId string) (*response.RoomBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getRoomURL(integrateId, roomId), headers.Header)

	result := new(response.RoomBean)
	err = utils.RespToBean(resp, result)

	return result, err
}
