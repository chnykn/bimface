package response

//OfflineDatabag ***
type OfflineDatabag struct {
	DatabagVersion string
	Status         string
	Reason         string
	CreateTime     string
}

//NewOfflineDatabag ***
func NewOfflineDatabag() *OfflineDatabag {
	o := &OfflineDatabag{
		// DatabagVersion:   ,
		// Status: ,
	}
	return o
}
