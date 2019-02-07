// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//ShareLink ***
type ShareLink struct {
	URL        string `json:"url"`
	ExpireTime string `json:"expireTime"`
}

//NewShareLink ***
func NewShareLink(url, expireTime string) *ShareLink {
	o := &ShareLink{
		URL:        url,
		ExpireTime: expireTime,
	}
	return o
}

// ToString get the string
func (o *ShareLink) ToString() string {
	return fmt.Sprintf("ShareLink [URL=%s, ExpireTime=%s]", o.URL, o.ExpireTime)
}
