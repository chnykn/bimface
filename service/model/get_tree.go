// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package model

import "fmt"

const (
	//获取构件分类树 POST https://api.bimface.com/data/v2/files/{fileId}/tree
	treeURI string = "/data/v2/files/%d/tree"
)

//---------------------------------------------------------------------

func (o *Service) treeURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+treeURI, fileId)
}

//---------------------------------------------------------------------
