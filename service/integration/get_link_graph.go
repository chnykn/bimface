// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	//获取集成文件链接关系 GET https://api.bimface.com/data/v2/integrations/{integrateId}/linkGraph
	linkGraphURI string = "/data/v2/integrations/%d/linkGraph"
)

func (o *Service) linkGraphURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+linkGraphURI, integrateId)
}

// 获取参与集成的子文件列表
func (o *Service) GetLinkGraphs(integrateId int64, includeDrawingSheet bool) ([]*response.LinkGraphNodeBean, error) {
	result := make([]*response.LinkGraphNodeBean, 0)
	err := o.GET(o.integrationFilesURL(integrateId, includeDrawingSheet), result)

	return result, err
}
