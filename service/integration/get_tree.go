// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"
	"strings"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/v2/bean/request"
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	//获取构件分类树 POST https://api.bimface.com/data/v2/integrations/{integrateId}/tree
	treeURI string = "/data/v2/integrations/%d/tree"
)

/* https://static.bimface.com/restful-apidoc/dist/modelIntegration.html#_gettreeusingpost_1

集成模型默认楼层分类树(v2.0), treeType接受两个值：floor, specialty和customized, 默认为floor.

当treeType为"customized"时
    desiredHierarchy表示了筛选树的层次,
    可选值有building,systemType,specialty,floor,category,family,familyType，
    如:desiredHierarchy=specialty,systemtype

    customizedNodeKeys: 用来指定筛选树每个维度用id或者是name作为唯一标识, 如"floor":"id"

	请求 body:
	{
	  "customizedNodeKeys" : {
		"string" : "string"
	  },
	  "desiredHierarchy" : [ "category", "family" ]
	}


https://api.bimface.com/data/v2/integrations/1211286711525184/tree
https://api.bimface.com/data/v2/integrations/1211286711525184/tree?treeType=floor
https://api.bimface.com/data/v2/integrations/1211286711525184/tree?treeType=specialty
https://api.bimface.com/data/v2/integrations/1211286711525184/tree?treeType=customized&desiredHierarchy=floor%2Cspecialty%2Ccategory%2Cfamily%2CfamilyType
*/

//---------------------------------------------------------------------

func (o *Service) treeURL(integrateId int64, treeType string, hierarchies []string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+treeURI, integrateId)

	if treeType == "floor" {
		result = result + "?treeType=floor"
	} else if treeType == "specialty" {
		result = result + "?treeType=specialty"
	} else if treeType == "customized" {
		result = result + "?treeType=customized"
		if len(hierarchies) > 0 {
			result = result + "&" + strings.Join(hierarchies, "%2C")
		}
	}

	return result
}

//---------------------------------------------------------------------

//获取构件分类树
func (o *Service) GetElementTree(integrateId int64, treeType string, hierarchies []string, treeRequest *request.IntegrationTreeOptionalRequest) ([]*response.ElementNodeBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	treeType = strings.ToLower(treeType)

	body := req.BodyJSON(treeRequest)
	resp := o.ServiceClient.Post(o.treeURL(integrateId, treeType, hierarchies), body, headers.Header)

	var result []*response.ElementNodeBean
	err = utils.RespToBean(resp, &result)

	return result, err
}
