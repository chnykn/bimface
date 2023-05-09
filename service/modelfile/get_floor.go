// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"
	"github.com/chnykn/bimface/v3/bean/response"
	"strconv"
	"strings"
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
	fileStrIds := make([]string, len(fileIds))
	for i := range fileIds {
		fileStrIds[i] = strconv.FormatInt(fileIds[i], 10)
	}

	fileIdsStr := strings.Join(fileStrIds, ",")
	return fmt.Sprintf(o.Endpoint.APIHost+floorsMappingsURI, fileIdsStr)
}

//---------------------------------------------------------------------

// 获取单模型的楼层信息
func (o *Service) GetFloors(fileId int64) ([]*response.FloorBean, error) {
	result := make([]*response.FloorBean, 0)
	err := o.GET(o.floorsURL(fileId), &result)

	return result, err
}

// 获取多个模型的楼层信息
func (o *Service) GetFilesFloors(fileIds []int64) (*response.FloorsBean, error) {
	result := new(response.FloorsBean)
	err := o.GET(o.floorsMappingsURL(fileIds), result)

	return result, err
}
