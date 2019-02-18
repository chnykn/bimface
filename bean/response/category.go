// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//Category ***
type Category struct {
	ID       string    `json:"categoryId"`
	Name     string    `json:"categoryName"`
	Families []*Family `json:"families"`
}

//NewCategory ***
func NewCategory() *Category { //id int64, name string
	o := &Category{
		// CategoryID:   id,
		// CategoryName: name,
		Families: make([]*Family, 0),
	}
	return o
}

// ToString get the string
func (o *Category) ToString() string {
	return fmt.Sprintf("CategoryBean [ID=%s, Name=%s, Families=%v]", o.ID, o.Name, o.Families)
}
