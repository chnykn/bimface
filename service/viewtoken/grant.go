// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package viewtoken

import "fmt"

const (
	viewTokenFileIdURI      string = "/view/token?fileId=%d"
	viewTokenIntegrateIdURI string = "/view/token?integrateId=%d"
	viewTokenCompareIdURI   string = "/view/token?compareId=%d"

	//获取子文件的View Token
	//https://api.bimface.com/data/v2/integrations/{integrateId}/files/{fileId}/viewToken
	viewTokenSubfileURI string = "/data/v2/integrations/%d/files/%d/viewToken"
)

//---------------------------------------------------------------------

func (o *Service) viewTokenFileIdURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+viewTokenFileIdURI, fileId)
}

func (o *Service) viewTokenIntegrateIdURL(integrateId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+viewTokenIntegrateIdURI, integrateId)
}

func (o *Service) viewTokenCompareIdURL(compareId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+viewTokenCompareIdURI, compareId)
}

func (o *Service) viewTokenSubfileURL(integrateId int64, fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+viewTokenSubfileURI, integrateId, fileId)
}

//---------------------------------------------------------------------

// 根据fileId、integrateId或者compareId获取viewToken，然后把viewToken传入JavaScript组件提供的接口中，即可显示工程文件。
// 注：只有在转换或集成任务成功以后，才能获取viewToken，有效期为12小时。
func (o *Service) grantViewTokenById(xxId int64, kind int) (string, error) {

	var url string
	switch kind {
	case 0:
		url = o.viewTokenFileIdURL(xxId)
	case 1:
		url = o.viewTokenIntegrateIdURL(xxId)
	case 2:
		url = o.viewTokenCompareIdURL(xxId)
	}

	if url == "" {
		return "", fmt.Errorf("url is null @ Service.grantViewTokenById")
	}

	result := new(string)
	err := o.GET(url, result)

	return *result, err
}

// 根据fileId获取viewToke，然后把viewToken传入JavaScript组件提供的接口中，即可显示工程文件。
// 注：只有在转换或集成任务成功以后，才能获取viewToken，有效期为12小时。
func (o *Service) GrantViewTokenByFileId(fileId int64) (string, error) {
	return o.grantViewTokenById(fileId, 0)
}

// 根据integrateId获取viewToke，然后把viewToken传入JavaScript组件提供的接口中，即可显示工程文件。
// 注：只有在转换或集成任务成功以后，才能获取viewToken，有效期为12小时。
func (o *Service) GrantViewTokenByIntegrateId(integrateId int64) (string, error) {
	return o.grantViewTokenById(integrateId, 1)
}

// 根据compareId获取viewToke，然后把viewToken传入JavaScript组件提供的接口中，即可显示工程文件。
// 注：只有在转换或集成任务成功以后，才能获取viewToken，有效期为12小时。
func (o *Service) GrantViewTokenByCompareId(compareId int64) (string, error) {
	return o.grantViewTokenById(compareId, 2)
}

// 获取子文件的View Token
// 注：一般使用在获取不随着模型集成而集成的模型信息，如图纸，视图等
func (o *Service) GrantSubfileViewToken(integrateId int64, fileId int64) (string, error) {

	url := o.viewTokenSubfileURL(integrateId, fileId)

	if url == "" {
		return "", fmt.Errorf("url is null @ Service.GrantSubfileViewToken")
	}

	result := new(string)
	err := o.GET(url, result)

	return *result, err
}
