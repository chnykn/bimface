// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sharelink

import (
	"fmt"
	"strconv"

	"github.com/chnykn/bimface/v2/bean/response"
	"github.com/chnykn/bimface/v2/utils"
)

const (
	//生成分享链接 POST https://api.bimface.com/share
	fileShareLinkURI      string = "/share?fileId=%d"      //&activeHours=%s
	integrateShareLinkURI string = "/share?integrateId=%d" //&activeHours=%s
)

// expireDate's format YYYY-MM-DD
func (o *Service) makeShareLinkURL(isFile bool, objectId int64, activeHours int, expireDate string, needPassword bool) string {

	var result string
	if isFile {
		result = fmt.Sprintf(o.Endpoint.APIHost+fileShareLinkURI, objectId)
	} else {
		result = fmt.Sprintf(o.Endpoint.APIHost+integrateShareLinkURI, objectId)
	}

	if activeHours > 0 {
		result = result + "&activeHours=" + strconv.Itoa(activeHours)
	} else if expireDate != "" {
		result = result + "&expireDate=" + expireDate
	}

	if needPassword {
		result = result + "&needPassword=true"
	}
	return result
}

//-------------------------------------------------------------

func (o *Service) doMakeShare(isFile bool, objectId int64, activeHours int, expireDate string, needPassword bool) (*response.ShareLinkBean, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	url := o.makeShareLinkURL(isFile, objectId, activeHours, expireDate, needPassword)
	resp := o.ServiceClient.Post(url, headers.Header)

	result := new(response.ShareLinkBean)
	err = utils.RespToBean(resp, result)

	return result, err
}

//生成模型文件的分享链接
func (o *Service) FileShareLink(fileId int64, activeHours int, expireDate string, needPassword bool) (*response.ShareLinkBean, error) {
	return o.doMakeShare(true, fileId, activeHours, expireDate, needPassword)
}

//生成模型集成的分享链接
func (o *Service) IntegrateShareLink(integrateId int64, activeHours int, expireDate string, needPassword bool) (*response.ShareLinkBean, error) {
	return o.doMakeShare(false, integrateId, activeHours, expireDate, needPassword)
}
