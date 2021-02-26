// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

//源文件删除

import (
	"fmt"

	"github.com/chnykn/bimface/bean"
	"github.com/chnykn/bimface/utils"
)

const (
	deleteFileURI string = "/file?fileId=%d"
)

//---------------------------------------------------------------------

func (o *Service) deleteFileURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+deleteFileURI, fileId)
}

//源文件相关: 删除文件
func (o *Service) DeleteFile(fileId int64) (string, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return "", err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Delete(o.deleteFileURL(fileId), headers.Header)

	var result *bean.RespResult
	result, err = utils.RespToResult(resp)

	var ret string
	if result != nil {
		ret = result.Code
	}

	return ret, nil
}
