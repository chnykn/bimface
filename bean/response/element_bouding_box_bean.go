package response

import "github.com/chnykn/bimface/v2/bean/common"

type ElementIdWithBoundingBox struct {
	ElementId   string             `json:"elementId"`
	BoundingBox common.BoundingBox `json:"boundingBox"`
}

type ElementBoundingBoxBean ElementIdWithBoundingBox
