package model

import (
	"fmt"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const (
	//获取单个面积分区信息 GET https://api.bimface.com/data/v2/files/{fileId}/areas/{areaId}
	getAreaURI string = "/data/v2/files/%d/areas/%d"

	//获取单个房间信息 GET https://api.bimface.com/data/v2/files/{fileId}/rooms/{roomId}
	getRoomURI string = "/data/v2/files/%d/rooms/%d"

	//获取楼层对应面积分区列表 GET https://api.bimface.com/data/v2/files/{fileId}/areas
	getAreasURI string = "data/v2/files/%d/areas"
)

//---------------------------------------------------------------------

func (o *Service) getAreaURL(fileId int64, areaId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getAreaURI, fileId, areaId)
}

func (o *Service) getRoomURL(fileId int64, roomId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getRoomURI, fileId, roomId)
}

func (o *Service) getAreasURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getAreasURI, fileId)
}

//------------------------------------------------------------------------------------

//获取单个面积分区信息
func (o *Service) GetArea(fileId int64, areaId int64) (*response.Area, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getAreaURL(fileId, areaId), headers.Header)

	var result *response.Area
	err = utils.RespToBean(resp, result)

	return result, err
}

//获取单个房间信息
func (o *Service) GetRoom(fileId int64, roomId int64) (*response.Area, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getRoomURL(fileId, roomId), headers.Header)

	var result *response.Area
	err = utils.RespToBean(resp, result)

	return result, err
}

//获取楼层对应面积分区列表
func (o *Service) GetAreas(fileId int64) ([]*response.Area, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getAreasURL(fileId), headers.Header)

	var result []*response.Area
	err = utils.RespToBean(resp, result)

	return result, err
}
