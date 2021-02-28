package integration

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/bean"
	"github.com/chnykn/bimface/bean/common"
	"github.com/chnykn/bimface/utils"
)

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

func (o *Service) modifyElementProperties(integrateId int64, fileId int64, elementId string, properties []*common.PropertyGroup, isDelete bool) (string, error) {
	if len(properties) <= 0 {
		return "", fmt.Errorf("properties is nil")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	body := req.BodyJSON(properties)
	var resp *req.Resp
	if isDelete {
		resp = o.ServiceClient.Delete(o.elementModifyPropertiesURL(integrateId, fileId, elementId), body, headers.Header)
	} else {
		resp = o.ServiceClient.Put(o.elementModifyPropertiesURL(integrateId, fileId, elementId), body, headers.Header)
	}

	var result *bean.RespResult
	result, err = utils.RespToResult(resp)

	var ret string
	if result != nil {
		ret = result.Code
	}

	return ret, nil
}

//修改单模型指定构件的属性
func (o *Service) SetElementProperties(integrateId int64, fileId int64, elementId string, properties []*common.PropertyGroup) (string, error) {
	return o.modifyElementProperties(integrateId, fileId, elementId, properties, false)
}

//删除单模型指定构件的属性
func (o *Service) DeleteElementProperties(integrateId int64, fileId int64, elementId string, properties []*common.PropertyGroup) (string, error) {
	return o.modifyElementProperties(integrateId, fileId, elementId, properties, true)
}
