package response

//FloorSpecialty Specialty in floor
type FloorSpecialty struct {
	Specialty  string
	Categories []Category
}

//NewFloorSpecialty ***
func NewFloorSpecialty(specialty string) *FloorSpecialty {
	o := &FloorSpecialty{
		Specialty:  specialty,
		Categories: make([]Category, 0),
	}
	return o
}

//---------------------------------------------------------------------

//Floor ***
type Floor struct {
	Floor       string
	Specialties []FloorSpecialty
}

//NewFloor ***
func NewFloor(floor string) *Floor {
	o := &Floor{
		Floor:       floor,
		Specialties: make([]FloorSpecialty, 0),
	}
	return o
}

//---------------------------------------------------------------------

//FloorTree ***
type FloorTree struct {
	TreeType int64
	Tree     []Floor
}

//NewFloorTree ***
func NewFloorTree(treeType int64) *FloorTree {
	o := &FloorTree{
		TreeType: treeType,
		Tree:     make([]Floor, 0),
	}
	return o
}
