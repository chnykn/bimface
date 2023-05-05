// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

import (
	"fmt"
)

// 获取文件下载地址
const (
	downloadURI string = "/download/url?fileId=%d" //&fileName=%s  fileName? name?
)

//---------------------------------------------------------------------

func (o *Service) downloadURL(fileId int64, fileName string) string {
	result := fmt.Sprintf(o.Endpoint.FileHost+downloadURI, fileId)
	if fileName != "" {
		result = result + "&name=" + fileName
	}
	return result
}

//---------------------------------------------------------------------

//源文件相关: 获取文件下载地址
/***
字段	类型	必填	描述
fileId	Number	Y	文件Id
name	String	N	自定义文件下载名
***/
func (o *Service) GetDownloadURL(fileId int64, fileName string) (string, error) {
	result := new(string)
	err := o.GET(o.downloadURL(fileId, fileName), result)

	return *result, err
}
