package response

//Integration ***
type Integration struct {
	IntegrateID int64
	Name        string
	Priority    int64
	Status      string
	Thumbnail   []string
	Reason      string
	CreateTime  string
}

//NewIntegration ***
func NewIntegration() *Integration { //id int64, name string
	o := &Integration{
		// ID:   id,
		// Name: name,
		Thumbnail: make([]string, 0),
	}
	return o
}
