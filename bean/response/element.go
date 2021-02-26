// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import (
	"fmt"

	"github.com/chnykn/bimface/bean/common"
)

//PropertyItem PropertyItem
type PropertyItem struct {
	Code      string `json:"code"`
	Extension string `json:"extension"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Unit      string `json:"unit,omitempty"`
	ValueType int    `json:"valueType,omitempty"`
}

// ToString get the string
func (o *PropertyItem) ToString() string {
	return fmt.Sprintf("PropertyItem [Key=%s, Value=%v, Unit=%s]", o.Key, o.Value, o.Unit)
}

//PropertyGroup PropertyGroup
type PropertyGroup struct {
	Group string          `json:"group"`
	Items []*PropertyItem `json:"items"`
}

//--------------------------------------------------------------------

//Element ***
type Element struct {
	Id          string             `json:"elementId"`
	Name        string             `json:"name"`
	GUId        string             `json:"guid"`
	FamilyGUId  string             `json:"familyGuid"`
	BoundingBox common.BoundingBox `json:"boundingBox"`
	Properties  []*PropertyGroup   `json:"properties"`
}

//AddPropertyGroup ***
func (o *Element) AddPropertyGroup(group *PropertyGroup) {
	o.Properties = append(o.Properties, group)
}

// ToString get the string
func (o *Element) ToString() string {
	return fmt.Sprintf("Element [FileId=%s, Name=%s, GUId=%s, FamilyGUId=%s,BoundingBox=%v, Properties=%v]",
		o.Id, o.Name, o.GUId, o.FamilyGUId, o.BoundingBox, o.Properties)
}
