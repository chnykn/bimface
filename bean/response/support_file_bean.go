// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

type FileSupportBean struct {
	Length int64    `json:"length"`
	Types  []string `json:"types"`
}

// ToString get the string
func (o *FileSupportBean) ToString() string {
	return fmt.Sprintf("FileSupportBean [Length=%d, Types=%v]", o.Length, o.Types)
}
