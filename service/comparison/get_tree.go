// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comparison

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/common"
)

const (
	//获取模型对比构件分类树 GET https://api.bimface.com/data/v2/comparisons/{comparisonId}/tree
	compareTreeURI string = "data/v2/comparisons/%d/tree"
)

func (o *Service) GetTree(compareId int64) (*common.Tree, error) {

	result := new(common.Tree)

	url := fmt.Sprintf(o.Endpoint.APIHost+compareTreeURI, compareId)
	err := o.GET(url, result)

	return result, err
}
