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

//NewIntgrFloorSpecialty ***
func NewIntgrFloorSpecialty(name string) *IntgrFloorSpecialty {
	o := &IntgrFloorSpecialty{
		Name:       name,
		Categories: make([]*Category, 0),
	}
	return o
}

//---------------------------------------------------------------------

//IntgrFloor ***
type IntgrFloor struct {
	Name        string                 `json:"floor"`
	Specialties []*IntgrFloorSpecialty `json:"specialties"`
}

//NewIntgrFloor ***
func NewIntgrFloor(name string) *IntgrFloor {
	o := &IntgrFloor{
		Name:        name,
		Specialties: make([]*IntgrFloorSpecialty, 0),
	}
	return o
}

//---------------------------------------------------------------------

//IntgrFloorTree ***
type IntgrFloorTree struct {
	TreeType int64         `json:"treeType"`
	Tree     []*IntgrFloor `json:"tree"`
}

//NewIntgrFloorTree ***
func NewIntgrFloorTree(treeType int64) *IntgrFloorTree {
	o := &IntgrFloorTree{
		TreeType: treeType,
		Tree:     make([]*IntgrFloor, 0),
	}
	return o
}
