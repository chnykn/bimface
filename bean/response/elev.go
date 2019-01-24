// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

/**
    "archElev": 7750,
    "areas": null,
    "elevation": 7750,
    "height": null,
    "id": "2",
    "miniMap": "m.bimface.com/ecf1e14463da72ddc43ed38fe1b7a4a7/resource/model/maps/3F/7.750.png",
    "name": "3F/7.750",
    "rooms": null,
	"structElev": 7750
**/

//Elev ***
type Elev struct {
	ID         string
	Name       string
	ArchElev   int64
	StructElev int64
	Elevation  int64
	Areas      string
	Height     string
	MiniMap    string //minimap picture httplink
	Rooms      string
}

//NewElev ***
func NewElev() *Elev { //id int64, name string
	o := &Elev{
		// ID:   id,
		// Name: name,
	}
	return o
}
