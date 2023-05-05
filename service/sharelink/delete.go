// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sharelink

import (
	"fmt"
	"strconv"

	"github.com/chnykn/bimface/v3/bean/response"
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

func (o *Service) doDeleteShareLink(isFile bool, objectId int64) error {
	url := o.deleteShareLinkURI(isFile, objectId)
	err := o.DELETE(url, nil)

	return err
}

// 取消模型文件的分享链接
func (o *Service) DeleteFileShareLink(fileId int64) error {
	return o.doDeleteShareLink(true, fileId)
}

// 取消模型集成的分享链接
func (o *Service) DeleteIntegrateShareLink(integrateId int64) error {
	return o.doDeleteShareLink(false, integrateId)
}

//--------------------------------------------------------------------------

func (o *Service) DeleteShareLinks(sourceIds []int64) (*response.BatchDeleteResultBean, error) {
	if len(sourceIds) <= 0 {
		return nil, fmt.Errorf("sourceIds is null")
	}

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
	url = url + sids

	result := new(response.BatchDeleteResultBean)
	err := o.DELETE(url, result)

	return result, err
}
