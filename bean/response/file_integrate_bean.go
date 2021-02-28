// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type FileIntegrateBean struct {
	IntegrateId int64  `json:"integrateId"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	SourceId    string `json:"string"`

	Status    string   `json:"status"`
	Reason    string   `json:"reason"`
	Thumbnail []string `json:"thumbnail"`

	CreateTime string `json:"createTime"`
}
