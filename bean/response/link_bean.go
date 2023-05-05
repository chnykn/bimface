// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type Link struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Guid      string `json:"guid"`
	Transform string `json:"transform"`
}

type LinkBean Link
