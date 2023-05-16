// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"github.com/chnykn/bimface/v3/bean/common"
	"github.com/chnykn/bimface/v3/service/comm"
)

/*
const (
	//修改单模型指定构件的属性
	//PUT https://api.bimface.com/data/v2/integrations/{integrateId}/files/{fileIdHash}/elements/{elementId}/properties

	//删除单模型指定构件的属性
	//DELETE https://api.bimface.com/data/v2/integrations/{integrateId}/files/{fileIdHash}/elements/{elementId}/properties

	elementModifyPropertiesURI string = "/data/v2/integrations/%d/files/%d/elements/%s/properties"
)

func (o *Service) elementModifyPropertiesURL(integrateId int64, fileId int64, elementId string) string {
	return fmt.Sprintf(o.Endpoint.APIHost+elementModifyPropertiesURI, integrateId, fileId, elementId)
}

func (o *Service) modifyElementProperties(integrateId int64, fileId int64, elementId string,
	properties []*common.PropertyGroup, isDelete bool) error {

	if len(properties) <= 0 {
		return fmt.Errorf("properties is nil")
	}

	var err error
	var body = httpkit.JsonReqBody(properties)

	if isDelete {
		err = o.DELETE(o.elementModifyPropertiesURL(integrateId, fileId, elementId), nil, body)
	} else {
		err = o.PUT(o.elementModifyPropertiesURL(integrateId, fileId, elementId), nil, body)
	}

	return err
}
*/

func (o *Service) modifyElementProperties(integrateId int64, fileId int64, elementId string,
	properties []*common.PropertyGroup, isDelete bool) error {
	return comm.ModifyElementProperties(o.Service, integrateId, fileId, elementId, properties, isDelete)
}

// 修改集成模型指定构件的属性
func (o *Service) SetElementProperties(integrateId int64, fileId int64, elementId string,
	properties []*common.PropertyGroup) error {

	return o.modifyElementProperties(integrateId, fileId, elementId, properties, false)
}

// 删除集成模型指定构件的属性
func (o *Service) DeleteElementProperties(integrateId int64, fileId int64, elementId string,
	properties []*common.PropertyGroup) error {

	return o.modifyElementProperties(integrateId, fileId, elementId, properties, true)
}
