// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package request

//ElementFilter ***
type GroupFilter struct {
	Group string   `json:"group"`
	Keys  []string `json:"keys,omitempty"`
}

//UploadRequest ***
type ElementFilterRequest struct {
	ElementIds []string      `json:"elementIds"`
	Filter     []GroupFilter `json:"filter,omitempty"`
}

//NewElementFilterRequest ***
func NewElementFilterRequest() *ElementFilterRequest {

	o := &ElementFilterRequest{
		ElementIds: make([]string, 0),
		Filter:     make([]GroupFilter, 0),
	}

	return o
}
