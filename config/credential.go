package config

// Credential ***
type Credential struct {
	AppKey    string
	AppSecret string
}

//NewCredential ***
func NewCredential(appKey string, appSecret string) *Credential {
	o := &Credential{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	return o
}
