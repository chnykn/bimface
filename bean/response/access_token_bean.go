// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

type AccessTokenBean struct {
	Token      string `json:"token"`
	ExpireTime string `json:"expireTime"`
}

func (o *AccessTokenBean) ToString() string {
	return fmt.Sprintf("AccessTokenBean [Token=%s, ExpireTime=%s]", o.Token, o.ExpireTime)
}
