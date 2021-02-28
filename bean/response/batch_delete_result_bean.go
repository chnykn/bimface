// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type BatchDeleteResultBean struct {
	Deleted      []int64 `json:"deleted"`      //
	Nonexistence []int64 `json:"nonexistence"` //不存在的
}
