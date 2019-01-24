package config

//Endpoint ***
type Endpoint struct {
	APIHost  string // consts.APIHost
	FileHost string // consts.FileHost
}

//NewEndpoint ***
func NewEndpoint(apiHost string, fileHost string) *Endpoint {
	o := &Endpoint{
		APIHost:  apiHost,
		FileHost: fileHost,
	}
	return o
}
