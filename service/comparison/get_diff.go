// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comparison

import (
	"fmt"
	"strconv"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	//分页获取模型对比结果 GET https://api.bimface.com/data/v2/comparisons/{compareId}/diff
	modelCompareDiffURI string = "data/v2/comparisons/%d/diff"

	//分页获取图纸对比结果 GET https://api.bimface.com/data/v2/comparisons/{compareId}/drawingdiff
	drawingCompareDiffURI string = "data/v2/comparisons/%d/drawingdiff"
)

//---------------------------------------------------------------------

func (o *Service) modelCompareDiffURL(compareId int64, family, elementName string, page, pageSize int) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+modelCompareDiffURI, compareId)

	s := "?"

	if family != "" {
		result = result + s + "family=?" + family
		s = "&"
	}

	if elementName != "" {
		result = result + s + "elementName=?" + elementName
		s = "&"
	}

	if page > 0 {
		result = result + s + "page=?" + strconv.Itoa(page)
		s = "&"
	}

	if pageSize > 0 {
		result = result + s + "pageSize=?" + strconv.Itoa(pageSize)
		s = "&"
	}

	return result
}

func (o *Service) drawingCompareDiffURL(compareId int64, layer string, page, pageSize int) string {
	result := fmt.Sprintf(o.Endpoint.APIHost+drawingCompareDiffURI, compareId)

	s := "?"

	if layer != "" {
		result = result + s + "layer=?" + layer
		s = "&"
	}

	if page > 0 {
		result = result + s + "page=?" + strconv.Itoa(page)
		s = "&"
	}

	if pageSize > 0 {
		result = result + s + "pageSize=?" + strconv.Itoa(pageSize)
		s = "&"
	}

	return result
}

//------------------------------------------------------------------------------------

// 分页获取模型对比结果  compareId必填，其他选填
func (o *Service) GetModelCompareDiff(compareId int64, family, elementName string,
	page, pageSize int) (*response.ModelCompareDiffsBean, error) {

	result := new(response.ModelCompareDiffsBean)

	url := o.modelCompareDiffURL(compareId, family, elementName, page, pageSize)
	err := o.GET(url, result)

	return result, err
}

// 分页获取图纸对比结果  compareId必填，其他选填
func (o *Service) GetDrawingCompareDiff(compareId int64, layer string,
	page, pageSize int) (*response.DrawingCompareDiffsBean, error) {

	result := new(response.DrawingCompareDiffsBean)

	url := o.drawingCompareDiffURL(compareId, layer, page, pageSize)
	err := o.GET(url, result)

	return result, err
}
