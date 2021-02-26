package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const (
	//获取单模型的楼层信息 GET https://api.bimface.com/data/v2/files/{fileId}/floors
	floorsURI string = "/data/v2/files/%d/floors"

	//获取多个模型的楼层信息 GET https://api.bimface.com/data/v2/files/{fileIds}/fileIdfloorsMappings
	floorsMappingsURI string = "/data/v2/files/%s/fileIdfloorsMappings"
)

//---------------------------------------------------------------------

func (o *Service) floorsURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+floorsURI, fileId)
}

func (o *Service) floorsMappingsURL(fileIds []int64) string {
	fileIDs := ""
	if len(fileIds) > 0 {
		for _, id := range fileIds {
			fileIDs = fileIDs + strconv.FormatInt(id, 10) + ","
		}
		fileIDs = strings.TrimRight(fileIDs, ",")
	}

	return fmt.Sprintf(o.Endpoint.APIHost+floorsMappingsURI, fileIDs)
}

//---------------------------------------------------------------------

//获取单模型的楼层信息
func (o *Service) GetFloors(fileId int64) ([]*response.Floor, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.floorsURL(fileId), headers.Header)

	var result []*response.Floor
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//获取多个模型的楼层信息
func (o *Service) GetFilesFloors(fileIds []int64) ([]*response.FileFloors, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.floorsMappingsURL(fileIds), headers.Header)

	var result []*response.FileFloors
	err = utils.RespToBean(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
