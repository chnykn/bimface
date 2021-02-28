// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comparison

import (
	"strconv"

	"github.com/chnykn/bimface/v2/bean"
	"github.com/chnykn/bimface/v2/utils"
)

//删除模型对比
func (o *Service) Delete(compareId int64) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	url := o.Endpoint.APIHost + compareURI + "?compareId=" + strconv.FormatInt(compareId, 10)
	resp := o.ServiceClient.Delete(url, headers.Header)

	var result *bean.RespResult
	result, err = utils.RespToResult(resp)

	var ret string
	if result != nil {
		ret = result.Code
	}

	return ret, nil
}
