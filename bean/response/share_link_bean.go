// Copyright 2019-2921 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/common"
)

type ShareLinkBean struct {
	URL        string `json:"url"`
	ExpireTime string `json:"expireTime"`
	Password   string `json:"password"`
	SourceId   string `json:"sourceId"`
	SourceName string `json:"sourceName"`
	SourceType string `json:"sourceType"`
	AppKey     string `json:"appKey"`
}

// ToString get the string
func (o *ShareLinkBean) ToString() string {
	return fmt.Sprintf("ShareLinkBean [URL=%s, ExpireTime=%s]", o.URL, o.ExpireTime)
}

//-----------------------------------------------------

type ShareLinkBeanPageList struct {
	common.PageList
	List []*ShareLinkBean `json:"list"`
}
