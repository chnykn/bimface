// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/httpkit"
)

const (
	//计算指定构件列表的包围盒 GET https://api.bimface.com/data/integrations/{integrateId}/elements/boundingboxes
	boundingBoxesURI string = "/data/integrations/%d/elements/boundingboxes"
)

func (o *Service) boundingBoxesURL(integrateId int64) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+boundingBoxesURI, integrateId)

	return result
}

//计算指定构件列表的包围盒
/* fileIdWithEleIdList: 构件ID列表(可选), 每个构件ID由fileID和elementID组成，两个Id之间用.隔开
[
	"1211224989204288.1572197",
	"1211224989204288.1572198"
]
*/
func (o *Service) GetBoundingBoxes(integrateId int64, fileIdWithEleIdList []string) ([]*response.ElementBoundingBoxBean, error) {

	var err error
	var result = make([]*response.ElementBoundingBoxBean, 0)

	if len(fileIdWithEleIdList) > 0 {
		body := httpkit.JsonReqBody(fileIdWithEleIdList)
		err = o.GET(o.boundingBoxesURL(integrateId), result, body)
	} else {
		err = o.GET(o.boundingBoxesURL(integrateId), result)
	}

	return result, err
}
