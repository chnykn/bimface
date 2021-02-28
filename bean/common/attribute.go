package common

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type ChangeAttribute struct {
	A *Attribute `json:"_A"`
	B *Attribute `json:"_B"`
}
