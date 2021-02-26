// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//"Intgr" stand for "Integration"

//IntgrSpecialtyFloor ***
type IntgrSpecialtyFloor struct {
	Name       string      `json:"floor"`
	Categories []*Category `json:"categories"`
}

//IntgrSpecialty ***
type IntgrSpecialty struct {
	Name   string                 `json:"specialty"`
	Floors []*IntgrSpecialtyFloor `json:"floors"`
}

//IntgrSpecialtyTree ***
type IntgrSpecialtyTree struct {
	TreeType int64             `json:"treeType"`
	Tree     []*IntgrSpecialty `json:"tree"`
}
