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

//Floor ***
type Floor struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Elevation  float64 `json:"elevation"`
	ArchElev   float64 `json:"archElev"`
	StructElev float64 `json:"structElev"`
	MiniMap    string  `json:"miniMap"`
	Areas      string  `json:"areas,omitempty"`
	Height     string  `json:"height,omitempty"`
	Rooms      string  `json:"rooms,omitempty"`
}

//NewFloor ***
func NewFloor() *Floor { //id int64, name string
	o := &Floor{
		// ID:   id,
		// Name: name,
	}
	return o
}
