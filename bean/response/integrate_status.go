// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//IntegrateStatus ***
type IntegrateStatus struct {
	IntegrateID int64
	Name        string
	Priority    int64
	Status      string
	Thumbnail   []string
	Reason      string
	CreateTime  string
}

//NewIntegrateStatus ***
func NewIntegrateStatus() *IntegrateStatus { //id int64, name string
	o := &IntegrateStatus{
		// ID:   id,
		// Name: name,
		Thumbnail: make([]string, 0),
	}
	return o
}
