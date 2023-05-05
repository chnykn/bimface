// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comparison

import (
	"strconv"
)

// 删除模型对比
func (o *Service) Delete(compareId int64) error {
	url := o.Endpoint.APIHost + compareURI + "?compareId=" + strconv.FormatInt(compareId, 10)
	err := o.DELETE(url, nil)

	return err
}
