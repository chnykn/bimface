// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

type GroupAndKeysPair struct {
	Group string   `json:"group"`
	Keys  []string `json:"keys,omitempty"`
}

type PropertyFilterRequest struct {
	ElementIds []string           `json:"elementIds"`
	Filter     []GroupAndKeysPair `json:"filter,omitempty"`
}

func NewPropertyFilterRequest() *PropertyFilterRequest {

	o := &PropertyFilterRequest{
		ElementIds: make([]string, 0),
		Filter:     make([]GroupAndKeysPair, 0),
	}

	return o
}
