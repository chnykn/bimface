package comparison

import (
	"fmt"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const (
	//获取模型构件对比差异 GET https://api.bimface.com/data/v2/comparisons/{comparisonId}/elementChange
	elementChangeURI string = "data/v2/comparisons/%d/elementChange?previousFileId=%d&previousElementId=%s&followingFileId=%d&followingElementId=%s"
)

func (o *Service) GetElementChange(compareId int64, previousFileId int64, previousElementId string,
	followingFileId int64, followingElementId string) (*response.ModelCompareChangeBean, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	url := fmt.Sprintf(o.Endpoint.APIHost+elementChangeURI, compareId, previousFileId, previousElementId, followingFileId, followingElementId)
	resp := o.ServiceClient.Get(url, headers.Header)

	var result *response.ModelCompareChangeBean
	err = utils.RespToBean(resp, result)

	return result, err
}
