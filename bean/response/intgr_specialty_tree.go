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

//NewIntgrSpecialtyFloor ***
func NewIntgrSpecialtyFloor(name string) *IntgrSpecialtyFloor {
	o := &IntgrSpecialtyFloor{
		Name:       name,
		Categories: make([]*Category, 0),
	}
	return o
}

//---------------------------------------------------------------------

//IntgrSpecialty ***
type IntgrSpecialty struct {
	Name   string                 `json:"specialty"`
	Floors []*IntgrSpecialtyFloor `json:"floors"`
}

//NewIntgrSpecialty ***
func NewIntgrSpecialty(name string) *IntgrSpecialty {
	o := &IntgrSpecialty{
		Name:   name,
		Floors: make([]*IntgrSpecialtyFloor, 0),
	}
	return o
}

//---------------------------------------------------------------------

//IntgrSpecialtyTree ***
type IntgrSpecialtyTree struct {
	TreeType int64             `json:"treeType"`
	Tree     []*IntgrSpecialty `json:"tree"`
}

//NewIntgrSpecialtyTree ***
func NewIntgrSpecialtyTree(treeType int64) *IntgrSpecialtyTree {
	o := &IntgrSpecialtyTree{
		TreeType: treeType,
		Tree:     make([]*IntgrSpecialty, 0),
	}
	return o
}
