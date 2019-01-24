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
	//获取文件转换的构件层次结构
	categoryURI string = "/data/hierarchy?fileId=%d"

	//获取集成模型的构件层次结构
	integrationTreeURI string = "/data/integration/tree?integrateId=%d&treeType=%d"
)

//CategoryTreeService ***
type CategoryTreeService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewCategoryTreeService ***
func NewCategoryTreeService(serviceClient *http.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *CategoryTreeService {
	o := &CategoryTreeService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //http.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *CategoryTreeService) categoryURL(fileID int64, isV2 bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+categoryURI, fileID)
	if isV2 {
		result = result + "&v=2.0"
	}
	return result
}

func (o *CategoryTreeService) integrationTreeURL(integrateID int64, treeType int) string {
	return fmt.Sprintf(o.Endpoint.APIHost+integrationTreeURI, integrateID, treeType)
}

//---------------------------------------------------------------------

//文件转换相关: 获取单文件的所有构件类别、族和族类型树
//http://static.bimface.com/book/restful/articles/api/translate/get-hierarchy.html
//1）获取1.0版本结果数据； 2）获取2.0版本结果数据
/***
字段	类型	必填	描述
fileId	Number	Y	文件ID
v		String	N	结果数据版本：1.0（1.0版本结果数据）2.0（2.0版本结果数据）	 默认为1.0数据
***/
func (o *CategoryTreeService) getCategoryTree(fileID int64, isV2 bool) (*req.Resp, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.categoryURL(fileID, isV2), headers.Header)
	return resp, nil
}

//GetCategoryTree 文件转换相关: 获取单文件的所有构件类别、族和族类型树, 结果数据版本：1.0（1.0版本结果数据）
/***
字段	类型	必填	描述
fileId	Number	Y	文件ID
***/
func (o *CategoryTreeService) GetCategoryTree(fileID int64) ([]response.Category, *utils.Error) {
	resp, err := o.getCategoryTree(fileID, false)
	if err != nil {
		return nil, err
	}

	result, err := http.RespToBeans(resp, &response.Category{})
	return result.([]response.Category), nil
}

//GetCategoryTreeV2 文件转换相关: 获取单文件的所有构件类别、族和族类型树, 结果数据版本：2.0（2.0版本结果数据）
/***
字段	类型	必填	描述
fileId	Number	Y	文件ID
***/
func (o *CategoryTreeService) GetCategoryTreeV2(fileID int64) ([]response.CategoryNode, *utils.Error) {
	resp, err := o.getCategoryTree(fileID, true)
	if err != nil {
		return nil, err
	}

	result, err := http.RespToBeans(resp, &response.CategoryNode{})
	return result.([]response.CategoryNode), nil
}

//---------------------------------------------------------------------

//模型集成相关: 获取集成模型的构件层次结构
//模型集成以后，可以获取两种构件的层次结构：1）按专业视图；2）按楼层视图
//http://static.bimface.com/book/restful/articles/api/integrate/get-integrate-tree.html
/***
字段		类型	必填	描述
integrateId	Number	Y	集成ID
treeType	Number	Y	树类型：1（按专业视图）2（按楼层视图）
***/
func (o *CategoryTreeService) getIntegrationTreeResp(integrateID int64, treeType int) (*req.Resp, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integrationTreeURL(integrateID, treeType), headers.Header)
	return resp, nil
}

//GetSpecialtyTree 按专业视图,获取集成模型的构件层次结构
/***
字段		类型	必填	描述
integrateId	Number	Y	集成ID
***/
func (o *CategoryTreeService) GetSpecialtyTree(integrateID int64) (*response.SpecialtyTree, *utils.Error) {
	resp, err := o.getIntegrationTreeResp(integrateID, 1)
	if err != nil {
		return nil, err
	}

	result := response.NewSpecialtyTree(1)
	err = http.RespToBean(resp, result)

	return result, err
}

//GetIntegrationSpecialtyTree same to GetSpecialtyTree
func (o *CategoryTreeService) GetIntegrationSpecialtyTree(integrateID int64) (*response.SpecialtyTree, *utils.Error) {
	return o.GetSpecialtyTree(integrateID)
}

//GetFloorTree 按楼层视图,获取集成模型的构件层次结构
/***
字段		类型	必填	描述
integrateId	Number	Y	集成ID
***/
func (o *CategoryTreeService) GetFloorTree(integrateID int64) (*response.FloorTree, *utils.Error) {
	resp, err := o.getIntegrationTreeResp(integrateID, 2)
	if err != nil {
		return nil, err
	}

	result := response.NewFloorTree(2)
	err = http.RespToBean(resp, result)

	return result, err
}

//GetIntegrationFloorTree same to GetFloorTree
func (o *CategoryTreeService) GetIntegrationFloorTree(integrateID int64) (*response.FloorTree, *utils.Error) {
	return o.GetFloorTree(integrateID)
}
