package response

import "fmt"

//SupportFile ***
type SupportFile struct {
	Length int64
	Types  []string
}

//NewSupportFile ***
func NewSupportFile() *SupportFile { //length int64, types []string
	o := &SupportFile{
		// Length: length,
		Types: make([]string, 0),
	}
	return o
}

// ToString get the string
func (o *SupportFile) ToString() string {
	return fmt.Sprintf("SupportFile [Length=%d, Types=%v]", o.Length, o.Types)
}
