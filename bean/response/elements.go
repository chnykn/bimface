package response

import "bimface/bean/common"

//Element ***
type Element struct {
	FileID    string
	ElementID string
}

//NewElement ***
func NewElement(fileID string, elementID string) *Element {
	o := &Element{
		FileID:    fileID,
		ElementID: elementID,
	}
	return o
}

//---------------------------------------------------------------------

//Elements ***
type Elements struct {
	Elements    []Element
	BoundingBox common.BoundingBox
}

//NewElements ***
func NewElements() *Elements { //id int64, name string
	o := &Elements{
		Elements: make([]Element, 0),
		//BoundingBox: common.NewBoundingBox(),
	}
	return o
}
