// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modelfile

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	//获取模型链接信息 GET https://api.bimface.com/data/v2/files/{fileId}/links
	linksURI string = "/data/v2/files/%d/links"
)

//---------------------------------------------------------------------

func (o *Service) linksURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+linksURI, fileId)
}

//---------------------------------------------------------------------

// 获取单模型的楼层信息
func (o *Service) GetLinks(fileId int64) ([]*response.LinkBean, error) {
	result := make([]*response.LinkBean, 0)
	err := o.GET(o.linksURL(fileId), &result)

	return result, err
}
