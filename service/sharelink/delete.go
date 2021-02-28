// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sharelink

import (
	"fmt"
	"strconv"

	"github.com/chnykn/bimface/v2/bean"
	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

func (o *Service) deleteShareLinkURI(isFile bool, objectId int64) string {

	var result string
	if isFile {
		result = fmt.Sprintf(o.Endpoint.APIHost+fileShareLinkURI, objectId)
	} else {
		result = fmt.Sprintf(o.Endpoint.APIHost+integrateShareLinkURI, objectId)
	}

	return result
}

//----------------------------------------------

func (o *Service) doDeleteShareLink(isFile bool, objectId int64) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	url := o.deleteShareLinkURI(isFile, objectId)
	resp := o.ServiceClient.Delete(url, headers.Header)

	var result *bean.RespResult
	result, err = utils.RespToResult(resp)

	var ret string
	if result != nil {
		ret = result.Code
	}

	return ret, nil
}

//取消模型文件的分享链接
func (o *Service) DeleteFileShareLink(fileId int64) (string, error) {
	return o.doDeleteShareLink(true, fileId)
}

//取消模型集成的分享链接
func (o *Service) DeleteIntegrateShareLink(integrateId int64) (string, error) {
	return o.doDeleteShareLink(false, integrateId)
}

//--------------------------------------------------------------------------

func (o *Service) DeleteShareLinks(sourceIds []int64) (*response.BatchDeleteResultBean, error) {
	if len(sourceIds) <= 0 {
		return nil, fmt.Errorf("sourceIds is null")
	}

	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	//--------------

	url := o.Endpoint.APIHost + shareLinksURI + "?sourceIds="

	var sids string
	high := len(sourceIds) - 1
	for i := range sourceIds {
		sids = sids + strconv.FormatInt(sourceIds[i], 10)
		if i < high {
			sids = sids + "%2C"
		}
	}
	sids = sids

	url = url + sids
	resp := o.ServiceClient.Delete(url, headers.Header)

	//--------------

	var result *response.BatchDeleteResultBean
	err = utils.RespToBean(resp, result)

	return result, err
}
