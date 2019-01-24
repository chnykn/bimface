package response

//CompareStatus ***
type CompareStatus struct {
	CompareID  int64
	Priority   int64
	Status     string
	Thumbnail  string // thumbnail http links
	Reason     string
	CreateTime string
}

//NewCompareStatus ***
func NewCompareStatus() *CompareStatus { //fileID int64, name string
	o := &CompareStatus{
		// FileID:   fileID,
		// Name: name,
		//Thumbnail: make([]string, 0),
	}
	return o
}

//=================================================================

//CompareElement ***
type CompareElement struct {
	DiffType           string // NEW|DELETE|CHANGE,
	Name               string // 柱1
	PreviousFileID     int64
	PreviousElementID  int64
	FollowingFileID    int64
	FollowingElementID int64
}

//----------------------------------------------------------

//CompareCategory ***
type CompareCategory struct {
	CategoryID   string
	CategoryName string
	ItemCount    int64
	Elements     []CompareElement
}

//NewCompareCategory ***
func NewCompareCategory() *CompareCategory {
	o := &CompareCategory{
		// CategoryID:   ,
		// CategoryName: ,
		Elements: make([]CompareElement, 0),
	}
	return o
}

//----------------------------------------------------------

//CompareData ***
type CompareData struct {
	SpecialtyID   string
	SpecialtyName string
	ItemCount     int64
	Categories    []CompareCategory
}

//NewCompareData ***
func NewCompareData() *CompareData {
	o := &CompareData{
		// SpecialtyID:   ,
		// SpecialtyName: ,
		Categories: make([]CompareCategory, 0),
	}
	return o
}
