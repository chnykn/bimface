// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//OfflineDatabag ***
type OfflineDatabag struct {
	DatabagVersion string
	Status         string
	Reason         string
	CreateTime     string
}

//NewOfflineDatabag ***
func NewOfflineDatabag() *OfflineDatabag {
	o := &OfflineDatabag{
		// DatabagVersion:   ,
		// Status: ,
	}
	return o
}
