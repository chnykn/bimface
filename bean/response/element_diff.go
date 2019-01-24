package response

//ElementAttr ***
type ElementAttr struct {
	Key   string //属性名
	Value string //属性值
	Unit  string //单位
}

//ChangeAttrs ***
type ChangeAttrs struct {
	A ElementAttr
	B ElementAttr
}

// ElementQty ***
type ElementQty struct {
	Code string  //工程量编码(用于运算)
	Desc string  //工程量描述
	Name string  //工程量名称(用于显示)
	Unit string  //单位
	Qty  float64 //数值
}

//ChangeQtys ***
type ChangeQtys struct {
	A ElementQty
	B ElementQty
}

//ElementDiff ***
type ElementDiff struct {
	A string //变化图元前一个版本的ID
	B string //变化图元后一个版本的ID

	NewAttributes    []ElementAttr
	DeleteAttributes []ElementAttr
	ChangeAttributes []ChangeAttrs

	NewQuantities    []ElementQty
	DeleteQuantities []ElementQty
	ChangeQuantities []ChangeQtys
}

//NewElementDiff ***
func NewElementDiff() *ElementDiff {
	o := &ElementDiff{
		NewAttributes:    make([]ElementAttr, 0),
		DeleteAttributes: make([]ElementAttr, 0),
		ChangeAttributes: make([]ChangeAttrs, 0),

		NewQuantities:    make([]ElementQty, 0),
		DeleteQuantities: make([]ElementQty, 0),
		ChangeQuantities: make([]ChangeQtys, 0),
	}
	return o
}
