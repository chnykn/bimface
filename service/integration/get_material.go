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

	//获取指定构件的材质信息
	//GET https://api.bimface.com/data/v2/integrations/{integrateId}/files/{fileIdHash}/elements/{elementId}/materials
	elementMaterialsURI string = "/data/v2/integrations/%d/files/%d/elements/%s/materials"
)

func (o *Service) elementMaterialsURL(integrateId int64, fileId int64, elementId string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+elementMaterialsURI, fileId, elementId)
}

//获取构件材质列表
func (o *Service) GetElementMaterials(integrateId int64, fileId int64, elementId string) ([]*response.MaterialBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.elementMaterialsURL(integrateId, fileId, elementId), headers.Header)

	var result []*response.MaterialBean
	err = utils.RespToBean(resp, &result)

	return result, err
}
