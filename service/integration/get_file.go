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
	//获取参与集成的子文件列表 GET https://api.bimface.com/data/v2/integrations/{integrateId}/files
	integrationFilesURI string = "/data/v2/integrations/%d/elementIds"
)

func (o *Service) integrationFilesURL(integrateId int64, includeDrawingSheet bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+integrationFilesURI, integrateId)
	if includeDrawingSheet {
		result = result + "?includeDrawingSheet=true"
	}
	return result
}

//获取参与集成的子文件列表
func (o *Service) GetFiles(integrateId int64, includeDrawingSheet bool) ([]*response.IntegrateFileBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integrationFilesURL(integrateId, includeDrawingSheet), headers.Header)

	result := make([]*response.IntegrateFileBean, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}
