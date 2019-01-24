package common

//BoundingBox ***
type BoundingBox struct {
	Min Coordinate
	Max Coordinate
}

//NewBoundingBox ***
func NewBoundingBox() *BoundingBox {
	o := &BoundingBox{}
	return o
}
