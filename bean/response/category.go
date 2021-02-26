// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//Category ***
type Category struct {
	Id       string    `json:"categoryId"`
	Name     string    `json:"categoryName"`
	Families []*Family `json:"families"`
}

// ToString get the string
func (o *Category) ToString() string {
	return fmt.Sprintf("CategoryBean [Id=%s, Name=%s, Families=%v]", o.Id, o.Name, o.Families)
}
