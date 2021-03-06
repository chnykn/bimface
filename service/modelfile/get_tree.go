// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"
	"strings"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/v2/bean/request"
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	//获取构件分类树 POST https://api.bimface.com/data/v2/files/{fileId}/tree
	//https://api.bimface.com/data/v2/files/1211223382064960/tree
	//https://api.bimface.com/data/v2/files/1211223382064960/tree?v=2.0
	//https://api.bimface.com/data/v2/files/1211223382064960/tree?v=2.0&treeType=customized
	treeURI string = "/data/v2/files/%d/tree?v=2.0"
)

/* https://static.bimface.com/restful-apidoc/dist/translateSingleModel.html#_gettreeusingpost

单模型构件分类树, treeType接受两个值：default和customized,默认为default.

v参数用来区别treeType为default时返回树的格式, customized总是返回格式2.0的构件树.

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


请求2.0的默认分类树：floor, category, family, familyType
https://api.bimface.com/data/v2/files/1211223382064960/tree?v=2.0

*/

//---------------------------------------------------------------------

func (o *Service) treeURL(fileId int64, treeType string) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+treeURI, fileId)

	if treeType == "customized" {
		result = result + "treeType=customized"
	}

	return result
}

//---------------------------------------------------------------------

//获取构件分类树
func (o *Service) GetElementTree(fileId int64, treeType string, treeRequest *request.FileTreeRequest) ([]*response.ElementNodeBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	treeType = strings.ToLower(treeType)

	var resp *req.Resp
	if (treeType == "customized") && (treeRequest != nil) {
		body := req.BodyJSON(treeRequest)
		resp = o.ServiceClient.Post(o.treeURL(fileId, treeType), body, headers.Header)
	} else {
		resp = o.ServiceClient.Post(o.treeURL(fileId, treeType), headers.Header)
	}

	result := make([]*response.ElementNodeBean, 0)
	err = utils.RespToBean(resp, &result)

	return result, err
}
