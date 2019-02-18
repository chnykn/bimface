// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//Family ***
type Family struct {
	ID    string   `json:"id"`
	Name  string   `json:"family"`
	Types []string `json:"familyTypes"`

	GUID string `json:"-"`
}

//NewFamily ***
func NewFamily() *Family {
	o := &Family{
		//Family:      name,
		Types: make([]string, 0),
	}
	return o
}

// ToString get the string
func (o *Family) ToString() string {
	return fmt.Sprintf("Family [Name=%s, Types=%v]", o.Name, o.Types)
}
