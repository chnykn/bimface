// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/common"
)

type Property struct {
	Id          string                  `json:"elementId"`
	Name        string                  `json:"name"`
	GUId        string                  `json:"guid"`
	FamilyGUId  string                  `json:"familyGuid"`
	BoundingBox common.BoundingBox      `json:"boundingBox"`
	Properties  []*common.PropertyGroup `json:"properties"`
}

func (o *Property) AddPropertyGroup(group *common.PropertyGroup) {
	o.Properties = append(o.Properties, group)
}

func (o *Property) ToString() string {
	return fmt.Sprintf("PropertyBean [FileId=%s, Name=%s, GUId=%s, FamilyGUId=%s,BoundingBox=%v, PropertyBean=%v]",
		o.Id, o.Name, o.GUId, o.FamilyGUId, o.BoundingBox, o.Properties)
}

//---------------------

type PropertyBean Property
