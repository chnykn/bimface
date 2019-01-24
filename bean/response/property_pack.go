package response

import (
	"bimface/bean/common"
	"fmt"
)

//PropertyItem PropertyItem
type PropertyItem struct {
	Key       string
	Value     string
	Unit      string
	ValueType int
}

//NewPropertyItem ***
func NewPropertyItem(key string, value string, unit string) *PropertyItem {
	o := &PropertyItem{
		Key:   key,
		Value: value,
		Unit:  unit,
	}
	return o
}

// ToString get the string
func (o *PropertyItem) ToString() string {
	return fmt.Sprintf("PropertyItem [Key=%s, Value=%v, Unit=%s]", o.Key, o.Value, o.Unit)
}

//--------------------------------------------------------------------

//PropertyGroup PropertyGroup
type PropertyGroup struct {
	Group string
	Items []PropertyItem
}

//NewPropertyGroup ***
func NewPropertyGroup(name string) *PropertyGroup {
	o := &PropertyGroup{
		Group: name,
		Items: make([]PropertyItem, 0),
	}
	return o
}

//--------------------------------------------------------------------

//PropertyPack ***
type PropertyPack struct {
	ElementID   string
	Name        string
	GUID        string
	FamilyGUID  string
	BoundingBox common.BoundingBox
	Properties  []PropertyGroup
}

//NewPropertyPack ***
func NewPropertyPack(elementID string) *PropertyPack {
	o := &PropertyPack{
		ElementID: elementID,
		//Groups: make([]PropertyGroup, 0),
	}
	return o
}

//AddPropertyGroup ***
func (o *PropertyPack) AddPropertyGroup(group PropertyGroup) {
	o.Properties = append(o.Properties, group)
}

// ToString get the string
func (o *PropertyPack) ToString() string {
	return fmt.Sprintf("Properties [ElementID=%s, Name=%s, GUID=%s, FamilyGUID=%s,BoundingBox=%v, Properties=%v]",
		o.ElementID, o.Name, o.GUID, o.FamilyGUID, o.BoundingBox, o.Properties)
}
