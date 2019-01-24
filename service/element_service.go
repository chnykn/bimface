package service

import (
	"bimface/bean/response"
	"bimface/config"
	"bimface/http"
	"bimface/utils"
	"fmt"

	"github.com/imroc/req"
)

const (
	//获取文件转换的构件列表
	elementURI string = "/data/element/id?fileId=%d"

	//获取集成模型的构件列表
	integrateElementURI string = "/data/integration/element?integrateId=%d"
)

//ElementService ***
type ElementService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewElementService ***
func NewElementService(serviceClient *http.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *ElementService {
	o := &ElementService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //http.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *ElementService) elementURL(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+elementURI, fileID)
}

func (o *ElementService) integrateElementURL(integrateID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+integrateElementURI, integrateID)
}

//---------------------------------------------------------------------

//GetElementsWithParams 文件转换相关: 获取文件转换的构件列表
//必填参数: fileID
func (o *ElementService) GetElementsWithParams(fileID int64, params req.QueryParam) ([]string, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.elementURL(fileID), params, headers.Header)

	result, err := http.RespToBeans(resp, new(string))
	if err != nil {
		return nil, err
	}

	return result.([]string), nil
}

//GetElements 文件转换相关: 获取文件转换的构件列表
//http://doc.bimface.com/book/restful/articles/api/integrate/get-integrate-element.html
/***
字段		类型	必填	描述
integrateId	Number	Y	集成ID
specialty	String	N	专业
floor		String	N	楼层
categoryId	String	N	构件分类ID
family		String	N	族
familyType	String	N	族类型
***/
func (o *ElementService) GetElements(fileID int64, floor, specialty, categoryID,
	family, familyType string) ([]string, *utils.Error) {

	params := make(req.QueryParam)
	if floor != "" {
		params["floor"] = floor
	}
	if specialty != "" {
		params["specialty"] = specialty
	}
	if categoryID != "" {
		params["categoryId"] = categoryID
	}
	if family != "" {
		params["family"] = family
	}
	if familyType != "" {
		params["familyType"] = familyType
	}

	return o.GetElementsWithParams(fileID, params)
}

//---------------------------------------------------------------------

//GetIntegrationElementsWithParams ***
func (o *ElementService) GetIntegrationElementsWithParams(integrateID int64,
	params req.QueryParam) (*response.Elements, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integrateElementURL(integrateID), params, headers.Header)

	result := response.NewElements()
	err = http.RespToBean(resp, result)

	return result, err
}

//GetIntegrationElements 模型集成相关: 获取集成的构件列表
//http://doc.bimface.com/book/restful/articles/api/translate/get-ele-ids.html
/***
字段		类型	必填	描述
fileId		Number	Y	文件ID
specialty	String	N	专业
floor		String	N	楼层
categoryId	String	N	构件分类ID
family		String	N	族
familyType	String	N	族类型
***/
func (o *ElementService) GetIntegrationElements(fileID int64, floor, specialty, categoryID,
	family, familyType string) (*response.Elements, *utils.Error) {

	params := make(req.QueryParam)
	if floor != "" {
		params["floor"] = floor
	}
	if specialty != "" {
		params["specialty"] = specialty
	}
	if categoryID != "" {
		params["categoryId"] = categoryID
	}
	if family != "" {
		params["family"] = family
	}
	if familyType != "" {
		params["familyType"] = familyType
	}

	return o.GetIntegrationElementsWithParams(fileID, params)
}
