// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"fmt"
	"strings"

	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/httpkit"
)

const (
	//获取构件分类树 POST https://api.bimface.com/data/v2/integrations/{integrateId}/tree
	treeURI string = "/data/v2/integrations/%d/tree"
)

/* https://static.bimface.com/restful-apidoc/dist/modelIntegration.html#_gettreeusingpost_1

集成模型默认楼层分类树(v2.0), treeType接受三个值：floor, specialty和customized, 默认为floor.
  floor     : floor, specialty, category, family, familyType
  specialty : specialty, floor, category, family, familyType
  customized:
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
https://api.bimface.com/data/v2/integrations/1211286711525184/tree?treeType=customized&desiredHierarchy=specialty%2Ccategory%2Cfamily%2CfamilyType
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
			result = result + "&desiredHierarchy=" + strings.Join(hierarchies, "%2C")
		}
	}

	return result
}

//---------------------------------------------------------------------

/* 获取构件分类树
treeType 可选三个值：floor（楼层）, specialty（专业）和customized（自定义）。
hierarchies 表示了筛选树的层次，当treeType为"customized"时才生效, 如：desiredHierarchy=specialty,systemtype。
treeRequest 中的
  - fileIdElementIds		由文件ID及构件ID构成对象组成的列表
  - sortedNamesHierarchy	排序名称层级
  - sorts					分类树节点排列数组
  - customizedNodeKeys		自定义节点属性
*/

func (o *Service) GetElementTree(integrateId int64, treeType string, hierarchies []string,
	treeRequest *request.IntegrationTreeOptionalRequest) (*response.ElementNodeTree, error) {

	treeType = strings.ToLower(treeType)
	result := new(response.ElementNodeTree)

	body := httpkit.JsonReqBody(treeRequest)
	err := o.POST(o.treeURL(integrateId, treeType, hierarchies), result, body)

	return result, err
}
