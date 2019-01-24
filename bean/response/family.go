// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//Family ***
type Family struct {
	Family      string
	FamilyTypes []string
}

//NewFamily ***
func NewFamily() *Family {
	o := &Family{
		//Family:      name,
		FamilyTypes: make([]string, 0),
	}
	return o
}

// ToString get the string
func (o *Family) ToString() string {
	return fmt.Sprintf("Family [Family=%s, FamilyTypes=%v]", o.Family, o.FamilyTypes)
}
