// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	//获取构件材质列表 GET https://api.bimface.com/data/v2/files/{fileId}/elements/{elementId}/materials
	elementMaterialsURI string = "/data/v2/files/%d/elements/%s/materials"
)

func (o *Service) elementMaterialsURL(fileId int64, elementId string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+elementMaterialsURI, fileId, elementId)
}

// 获取构件材质列表
func (o *Service) GetElementMaterials(fileId int64, elementId string) ([]*response.MaterialBean, error) {
	result := make([]*response.MaterialBean, 0)
	err := o.GET(o.elementMaterialsURL(fileId, elementId), &result)

	return result, err
}
