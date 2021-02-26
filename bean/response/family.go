// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "fmt"

//Family ***
type Family struct {
	Id    string   `json:"id"`
	Name  string   `json:"family"`
	Types []string `json:"familyTypes"`

	GUId string `json:"-"`
}

// ToString get the string
func (o *Family) ToString() string {
	return fmt.Sprintf("Family [Name=%s, Types=%v]", o.Name, o.Types)
}
