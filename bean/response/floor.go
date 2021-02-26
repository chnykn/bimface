// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import "github.com/chnykn/bimface/bean/common"

/**
{
 	"id" : "311",
	"name" : "elevation 1",
    "archElev" : 0.0,
    "elevation" : 0.0,
    "structElev" : 0.0
    "height" : 4000.0,
    "miniMap" : "787e5907b0ca5cb35f5d10ba091a085b/resource/model/maps/elevation 1.png",

    "areas" : [ {
      "boundary" : "",
      "id" : "313137",
      "levelId" : "11",
      "maxPt" : {
        "x" : -4938.068482562385,
        "y" : -3201.59397858169,
        "z" : 0.0
      },
      "minPt" : {
        "x" : -4938.068482562385,
        "y" : -3201.59397858169,
        "z" : 0.0
      },
      "name" : "dining room 4"
    } ],
    "rooms" : [ {
      "boundary" : "",
      "id" : "313137",
      "levelId" : "11",
      "maxPt" : {
        "x" : -4938.068482562385,
        "y" : -3201.59397858169,
        "z" : 0.0
      },
      "minPt" : {
        "x" : -4938.068482562385,
        "y" : -3201.59397858169,
        "z" : 0.0
      },
      "name" : "dining room 4"
    } ]

  }
**/

type FloorArea struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Boundary string `json:"boundary"`
	LevelId  string `json:"levelId"`

	MaxPt *common.Coordinate `json:"maxPt,omitempty"`
	MinPt *common.Coordinate `json:"minPt,omitempty"`
}

//Floor ***
type Floor struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Elevation  float64 `json:"elevation"`
	ArchElev   float64 `json:"archElev"`
	StructElev float64 `json:"structElev"`
	Height     float64 `json:"height"`
	MiniMap    string  `json:"miniMap"`

	Areas []FloorArea `json:"areas,omitempty"`
	Rooms []FloorArea `json:"rooms,omitempty"`
}

//-----------------------

//FileFloors ***
type FileFloors struct {
	FileId int64   `json:"FileId"`
	Floors []Floor `json:"floors,omitempty"`
}
