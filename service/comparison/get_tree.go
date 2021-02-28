package comparison

import (
	"fmt"

	"github.com/chnykn/bimface/v2/bean/common"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	//获取模型对比构件分类树 GET https://api.bimface.com/data/v2/comparisons/{comparisonId}/tree
	compareTreeURI string = "data/v2/comparisons/%d/tree"
)

func (o *Service) GetTree(compareId int64) (*common.Tree, error) {

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	url := fmt.Sprintf(o.Endpoint.APIHost+compareTreeURI, compareId)
	resp := o.ServiceClient.Get(url, headers.Header)

	var result *common.Tree
	err = utils.RespToBean(resp, result)

	return result, err
}
