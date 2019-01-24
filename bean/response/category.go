package response

import "fmt"

//Category ***
type Category struct {
	CategoryID   string
	CategoryName string
	Families     []Family
}

//NewCategory ***
func NewCategory() *Category { //id int64, name string
	o := &Category{
		// CategoryID:   id,
		// CategoryName: name,
		Families: make([]Family, 0),
	}
	return o
}

// ToString get the string
func (o *Category) ToString() string {
	return fmt.Sprintf("CategoryBean [CategoryID=%s, CategoryName=%s, Families=%v]",
		o.CategoryID, o.CategoryName, o.Families)
}
