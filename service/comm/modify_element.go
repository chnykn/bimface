package comm

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/common"
	"github.com/chnykn/bimface/v3/service/abstract"
	"github.com/chnykn/httpkit"
)

const (
	//修改集成模型指定构件的属性
	//PUT https://api.bimface.com/data/v2/integrations/{integrateId}/files/{fileIdHash}/elements/{elementId}/properties

	//删除集成模型指定构件的属性
	//DELETE https://api.bimface.com/data/v2/integrations/{integrateId}/files/{fileIdHash}/elements/{elementId}/properties

	//修改模型文件指定构件的属性
	//PUT https://api.bimface.com/data/v2/files/{fileId}/elements/{elementId}/properties

	//删除单文件模型指定构件的属性
	//DELETE https://api.bimface.com/data/v2/files/{fileId}/elements/{elementId}/properties

	fileElementModifyPropertiesURI  string = "/data/v2/files/%d/elements/%s/properties"
	intgrElementModifyPropertiesURI string = "/data/v2/integrations/%d/files/%d/elements/%s/properties"
)

func elementModifyPropertiesURL(service *abstract.Service, integrateId int64, fileId int64, elementId string) string {
	if integrateId > 0 {
		return fmt.Sprintf(service.Endpoint.APIHost+intgrElementModifyPropertiesURI, integrateId, fileId, elementId)
	}
	return fmt.Sprintf(service.Endpoint.APIHost+fileElementModifyPropertiesURI, fileId, elementId)
}

func ModifyElementProperties(service *abstract.Service, integrateId int64, fileId int64, elementId string,
	properties []*common.PropertyGroup, isDelete bool) error {

	if len(properties) <= 0 {
		return fmt.Errorf("properties is nil")
	}

	var err error
	var body = httpkit.JsonReqBody(properties)

	if isDelete {
		err = service.DELETE(elementModifyPropertiesURL(service, integrateId, fileId, elementId), nil, body)
	} else {
		err = service.PUT(elementModifyPropertiesURL(service, integrateId, fileId, elementId), nil, body)
	}

	return err
}
