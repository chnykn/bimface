package common

type Quantity struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
	Unit string `json:"unit"`
	Desc string `json:"decs,omitempty"`
}

type ChangeQuantity struct {
	A *Quantity `json:"_A"`
	B *Quantity `json:"_B"`
}
