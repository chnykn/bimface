// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type ChangeAttribute struct {
	A *Attribute `json:"_A"`
	B *Attribute `json:"_B"`
}
