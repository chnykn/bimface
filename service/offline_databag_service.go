package service

import (
	"bimface/bean/request"
	"bimface/bean/response"
	"bimface/config"
	"bimface/http"
	"bimface/utils"
	"fmt"
)

/**
const (
	createOfflineDatabagByFileID      string = "/files/%d/offlineDatabag?callback=%s"
	createOfflineDatabagByIntegrateID string = "/integrations/%d/offlineDatabag?callback=%s"
	createOfflineDatabagByCompareID   string = "/comparisions/%d/offlineDatabag?callback=%s"

	queryOfflineDatabagByFileID      string = "/files/%d/offlineDatabag"
	queryOfflineDatabagByIntegrateID string = "/integrations/%d/offlineDatabag"
	queryOfflineDatabagByCompareID   string = "/comparisions/%d/offlineDatabag"

	getOfflineDatabagURLByFileID      string = "/data/databag/downloadUrl?fileId=%d&type=offline&databagVersion=%s"
	getOfflineDatabagURLByIntegrateID string = "/data/databag/downloadUrl?integrateId=%d&type=offline&databagVersion=%s"
	getOfflineDatabagURLByCompareID   string = "/data/databag/downloadUrl?comapreId=%d&type=offline&databagVersion=%s"
)
**/

const (
	createOfflineDatabagURI string = "/%s/%d/offlineDatabag" //?callback=%s
	queryOfflineDatabagURI  string = "/%s/%d/offlineDatabag"
	getOfflineDatabagURI    string = "/data/databag/downloadUrl?%s=%d&type=offline" //&databagVersion=%s
)

//OfflineDatabagService ***
type OfflineDatabagService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewOfflineDatabagService ***
func NewOfflineDatabagService(serviceClient *http.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *OfflineDatabagService {
	o := &OfflineDatabagService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //http.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------
// kind must in [files, integrations, comparisions]
func (o *OfflineDatabagService) createOfflineDatabagURL(kind string, xxID int64, callback string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+createOfflineDatabagURI, kind, xxID)
	if callback != "" {
		result = result + "?callback=" + http.EncodeURI(callback)
	}
	return result
}

func (o *OfflineDatabagService) queryOfflineDatabagURL(kind string, xxID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+queryOfflineDatabagURI, kind, xxID)
}

//databagVersion 数据包版本，如果只有一个，则下载唯一的数据包，如果多个，则必须指定数据包版本
func (o *OfflineDatabagService) downloadOfflineDatabagURL(kind string, xxID int64, databagVersion string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+getOfflineDatabagURI, kind, xxID)
	if databagVersion != "" {
		result = result + "&databagVersion=" + databagVersion
	}
	return result
}

//-------------------------------------------------------------------------------

//CreateOfflineDatabag 离线数据包相关: 创建离线数据包
//http://static.bimface.com/book/restful/articles/api/offlinedatabag/create-offlinedatabag.html
/***
字段		类型	必填	描述
fileId		Number	Y	通过文件转换ID创建离线数据包时必填
integrateId	Number	Y	通过集成模型ID创建离线数据包时必填
compareId	Number	Y	通过模型对比ID创建离线数据包时必填
callback	String	N	回调url
***/
func (o *OfflineDatabagService) CreateOfflineDatabag(databagRequest *request.OfflineDatabagRequest) (*response.OfflineDatabag, *utils.Error) {

	var url string
	if databagRequest.FileID != nil {
		url = o.createOfflineDatabagURL("files", *databagRequest.FileID, databagRequest.Callback)
	} else if databagRequest.IntegrateID != nil {
		url = o.createOfflineDatabagURL("integrations", *databagRequest.IntegrateID, databagRequest.Callback)
	} else if databagRequest.CompareID != nil {
		url = o.createOfflineDatabagURL("comparisions", *databagRequest.IntegrateID, databagRequest.Callback)
	}
	if url == "" {
		return nil, utils.NewError("url is null", "url = '' @ OfflineDatabagService.createOfflineDatabag")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Put(url, headers.Header)

	result := response.NewOfflineDatabag()
	err = http.RespToBean(resp, result)

	return result, err
}

//-------------------------------------------------------------------------------

//QueryOfflineDatabag 离线数据包相关: 查询离线数据包
//http://static.bimface.com/book/restful/articles/api/offlinedatabag/query-offlinedataba.html
/***
字段		类型	必填	描述
fileId		Number	Y	通过文件转换ID创建离线数据包时必填
integrateId	Number	Y	通过集成模型ID创建离线数据包时必填
compareId	Number	Y	通过模型对比ID创建离线数据包时必填
***/
func (o *OfflineDatabagService) QueryOfflineDatabag(databagRequest *request.OfflineDatabagRequest) ([]response.OfflineDatabag, *utils.Error) {

	var url string
	if databagRequest.FileID != nil {
		url = o.queryOfflineDatabagURL("files", *databagRequest.FileID)
	} else if databagRequest.IntegrateID != nil {
		url = o.queryOfflineDatabagURL("integrations", *databagRequest.IntegrateID)
	} else if databagRequest.CompareID != nil {
		url = o.queryOfflineDatabagURL("comparisions", *databagRequest.IntegrateID)
	}
	if url == "" {
		return nil, utils.NewError("url is null", "url = '' @ OfflineDatabagService.queryOfflineDatabag")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(url, headers.Header)

	result, err := http.RespToBeans(resp, &response.OfflineDatabag{})
	return result.([]response.OfflineDatabag), nil
}

//-------------------------------------------------------------------------------

//GetOfflineDatabagDownloadURL 离线数据包相关: 获取离线数据包下载地址
//http://static.bimface.com/book/restful/articles/api/offlinedatabag/get-download-offlinedataba-url.html
/***
字段			类型	必填	描述
fileId			Number	Y	通过文件转换ID获取离线数据包下载地址时必填
integrateId		Number	Y	通过集成模型ID获取离线数据包下载地址时必填
compareId		Number	Y	通过模型对比ID获取离线数据包下载地址时必填
type			String	Y	值必须是“offline”
databagVersion	String	N	数据包版本，如果只有一个，则下载唯一的数据包，如果多个，则必须指定数据包版本, 例如 3.0
***/
func (o *OfflineDatabagService) GetOfflineDatabagDownloadURL(databagRequest *request.OfflineDatabagRequest) (string, *utils.Error) {

	var url string
	if databagRequest.FileID != nil {
		url = o.downloadOfflineDatabagURL("fileId", *databagRequest.FileID, databagRequest.DatabagVersion)
	} else if databagRequest.IntegrateID != nil {
		url = o.downloadOfflineDatabagURL("integrateId", *databagRequest.IntegrateID, databagRequest.DatabagVersion)
	} else if databagRequest.CompareID != nil {
		url = o.downloadOfflineDatabagURL("comapreId", *databagRequest.IntegrateID, databagRequest.DatabagVersion)
	}
	if url == "" {
		return "", utils.NewError("url is null", "url = '' @ OfflineDatabagService.queryOfflineDatabag")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(url, headers.Header)

	result := new(string)
	err = http.RespToBean(resp, result)

	if err != nil {
		return "", err
	}

	return *result, nil
}
