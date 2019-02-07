// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//FloorSpecialty Specialty in floor
type FloorSpecialty struct {
	Specialty  string     `json:"specialty"`
	Categories []Category `json:"categories"`
}

//NewFloorSpecialty ***
func NewFloorSpecialty(specialty string) *FloorSpecialty {
	o := &FloorSpecialty{
		Specialty:  specialty,
		Categories: make([]Category, 0),
	}
	return o
}

//---------------------------------------------------------------------

//Floor ***
type Floor struct {
	Name        string           `json:"floor"`
	Specialties []FloorSpecialty `json:"specialties"`
}

//NewFloor ***
func NewFloor(name string) *Floor {
	o := &Floor{
		Name:        name,
		Specialties: make([]FloorSpecialty, 0),
	}
	return o
}

//---------------------------------------------------------------------

//FloorTree ***
type FloorTree struct {
	TreeType int64   `json:"treeType"`
	Tree     []Floor `json:"tree"`
}

//NewFloorTree ***
func NewFloorTree(treeType int64) *FloorTree {
	o := &FloorTree{
		TreeType: treeType,
		Tree:     make([]Floor, 0),
	}
	return o
}
