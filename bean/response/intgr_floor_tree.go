// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//"Intgr" stand for "Integration"

//IntgrFloorSpecialty Specialty in floor
type IntgrFloorSpecialty struct {
	Name       string      `json:"specialty"`
	Categories []*Category `json:"categories"`
}

//IntgrFloor ***
type IntgrFloor struct {
	Name        string                 `json:"floor"`
	Specialties []*IntgrFloorSpecialty `json:"specialties"`
}

//IntgrFloorTree ***
type IntgrFloorTree struct {
	TreeType int64         `json:"treeType"`
	Tree     []*IntgrFloor `json:"tree"`
}
