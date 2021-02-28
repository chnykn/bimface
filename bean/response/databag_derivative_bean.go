// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type DataBagDerivativeBean struct {
	Length         int64  `json:"length"`
	DataBagVersion string `json:"databagVersion"`

	Reason string `json:"reason"`
	Status string `json:"status"`

	CreateTime string `json:"createTime"`
}

type DataBagBean DataBagDerivativeBean
