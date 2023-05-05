// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	//获取指定构件的材质信息
	//GET https://api.bimface.com/data/v2/integrations/{integrateId}/files/{fileIdHash}/elements/{elementId}/materials
	elementMaterialsURI string = "/data/v2/integrations/%d/files/%d/elements/%s/materials"
)

func (o *Service) elementMaterialsURL(integrateId int64, fileId int64, elementId string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+elementMaterialsURI, integrateId, fileId, elementId)
}

// 获取构件材质列表
func (o *Service) GetElementMaterials(integrateId int64, fileId int64, elementId string) ([]*response.MaterialBean, error) {
	result := make([]*response.MaterialBean, 0)
	err := o.GET(o.elementMaterialsURL(integrateId, fileId, elementId), result)

	return result, err
}
