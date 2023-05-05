// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comparison

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	//获取模型构件对比差异 GET https://api.bimface.com/data/v2/comparisons/{comparisonId}/elementChange
	elementChangeURI string = "data/v2/comparisons/%d/elementChange?previousFileId=%d&previousElementId=%s&followingFileId=%d&followingElementId=%s"
)

func (o *Service) GetElementChange(compareId int64, previousFileId int64, previousElementId string,
	followingFileId int64, followingElementId string) (*response.ModelCompareChangeBean, error) {

	result := new(response.ModelCompareChangeBean)

	url := fmt.Sprintf(o.Endpoint.APIHost+elementChangeURI, compareId, previousFileId, previousElementId, followingFileId, followingElementId)
	err := o.GET(url, result)

	return result, err
}
