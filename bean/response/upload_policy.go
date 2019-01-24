package response

//UploadPolicy ***
type UploadPolicy struct {
	Host         string
	Policy       string
	AccessID     string
	Signature    string
	Expire       int64
	CallbackBody string
	ObjectKey    string
}

//NewUploadPolicy ***
func NewUploadPolicy() *UploadPolicy { //host string, policy string
	o := &UploadPolicy{
		// Host:   host,
		// Policy: policy,
	}
	return o
}
