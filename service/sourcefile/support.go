// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sourcefile

import (
	"github.com/chnykn/bimface/v3/bean/response"
)

const supportFileURI = "/support"

var defSupportFile = &response.FileSupportBean{
	Length: 50,
	Types: []string{
		"3dm", "3ds", "3dxml", "asm", "catpart", "catproduct", "dae", "dgn", "dwf", "dwfx",
		"dwg", "dxf", "fbx", "gbg", "gbq", "gcl", "gdq", "ggj", "gjg", "gmp",
		"gpb", "gpv", "gqi", "gsc", "gsh", "gtb", "gtj", "gzb", "iam", "ifc",
		"igms", "ipt", "jt", "nwc", "nwd", "obj", "osgb", "ply", "prt", "rfa",
		"rte", "rvm", "rvt", "shp", "skp", "sldasm", "sldprt", "step", "stl", "stp",
	},
}

//---------------------------------------------------------------------

func (o *Service) supportFileURL() string {
	return o.Endpoint.FileHost + supportFileURI
}

// GetSupportFile ***
func (o *Service) GetSupportFile() (*response.FileSupportBean, error) {

	if o.supportFile != nil {
		return o.supportFile, nil
	}

	//---------------------

	result := new(response.FileSupportBean)
	err := o.GET(o.supportFileURL(), result)

	if err == nil {
		o.supportFile = result
	} else {
		o.supportFile = defSupportFile
	}

	return result, err
}
