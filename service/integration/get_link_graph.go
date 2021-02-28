package integration

import (
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
	"fmt"
)

const (
	//获取集成文件链接关系 GET https://api.bimface.com/data/v2/integrations/{integrateId}/linkGraph
	linkGraphURI string = "/data/v2/integrations/%d/linkGraph"
)

func (o *Service) linkGraphURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+linkGraphURI, integrateId)
}

//获取参与集成的子文件列表
func (o *Service) GetLinkGraphs(integrateId int64, includeDrawingSheet bool) ([]*response.LinkGraphNodeBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.integrationFilesURL(integrateId, includeDrawingSheet), headers.Header)

	var result []*response.LinkGraphNodeBean
	err = utils.RespToBean(resp, &result)

	return result, err
}
