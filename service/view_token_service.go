// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/utils"
)

const (
	viewTokenFileIDURI      string = "/view/token?fileId=%d"
	viewTokenIntegrateIDURI string = "/view/token?integrateId=%d"
	viewTokenCompareIDURI   string = "/view/token?compareId=%d"
)

//ViewTokenService ***
type ViewTokenService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewViewTokenService ***
func NewViewTokenService(serviceClient *utils.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *ViewTokenService {
	o := &ViewTokenService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //utils.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *ViewTokenService) viewTokenFileIDURL(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+viewTokenFileIDURI, fileID)
}

func (o *ViewTokenService) viewTokenIntegrateIDURL(integrateID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+viewTokenIntegrateIDURI, integrateID)
}

func (o *ViewTokenService) viewTokenCompareIDURL(compareID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+viewTokenCompareIDURI, compareID)
}

//---------------------------------------------------------------------

//根据fileId、integrateId或者compareId获取viewToken，然后把viewToken传入JavaScript组件提供的接口中，即可显示工程文件。
//注：只有在转换或集成任务成功以后，才能获取viewToken，有效期为12小时。
func (o *ViewTokenService) grantViewTokenByID(xxID int64, kind int) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	var url string
	switch kind {
	case 0:
		url = o.viewTokenFileIDURL(xxID)
	case 1:
		url = o.viewTokenIntegrateIDURL(xxID)
	case 2:
		url = o.viewTokenCompareIDURL(xxID)
	}

	if url == "" {
		return "", fmt.Errorf("url is null @ ViewTokenService.grantViewTokenByID")
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(url, headers.Header)

	result := new(string)
	err = utils.RespToBean(resp, result)

	return *result, err
}

//GrantViewTokenByFileID 根据fileId获取viewToke，然后把viewToken传入JavaScript组件提供的接口中，即可显示工程文件。
//注：只有在转换或集成任务成功以后，才能获取viewToken，有效期为12小时。
func (o *ViewTokenService) GrantViewTokenByFileID(fileID int64) (string, error) {
	return o.grantViewTokenByID(fileID, 0)
}

//GrantViewTokenByIntegrateID 根据integrateId获取viewToke，然后把viewToken传入JavaScript组件提供的接口中，即可显示工程文件。
//注：只有在转换或集成任务成功以后，才能获取viewToken，有效期为12小时。
func (o *ViewTokenService) GrantViewTokenByIntegrateID(integrateID int64) (string, error) {
	return o.grantViewTokenByID(integrateID, 1)
}

//GrantViewTokenByCompareID 根据compareId获取viewToke，然后把viewToken传入JavaScript组件提供的接口中，即可显示工程文件。
//注：只有在转换或集成任务成功以后，才能获取viewToken，有效期为12小时。
func (o *ViewTokenService) GrantViewTokenByCompareID(compareID int64) (string, error) {
	return o.grantViewTokenByID(compareID, 2)
}
