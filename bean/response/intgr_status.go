// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//"Intgr" stand for "Integration"

//IntgrStatus ***
type IntgrStatus struct {
	ID         int64    `json:"integrateId"`
	Name       string   `json:"name"`
	Priority   int64    `json:"priority"`
	Status     string   `json:"status"`
	Thumbnail  []string `json:"thumbnail"`
	Reason     string   `json:"reason"`
	CreateTime string   `json:"createTime"`
}

//NewIntgrStatus ***
func NewIntgrStatus() *IntgrStatus { //id int64, name string
	o := &IntgrStatus{
		// ID:   id,
		// Name: name,
		Thumbnail: make([]string, 0),
	}
	return o
}
