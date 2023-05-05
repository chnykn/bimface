// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

//源文件删除

import (
	"fmt"
)

const (
	deleteFileURI string = "/file?fileId=%d"
)

//---------------------------------------------------------------------

func (o *Service) deleteFileURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+deleteFileURI, fileId)
}

// 源文件相关: 删除文件
func (o *Service) Delete(fileId int64) error {
	err := o.DELETE(o.deleteFileURL(fileId), nil)
	return err
}
