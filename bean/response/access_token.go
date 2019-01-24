package response

import "fmt"

//AccessToken ***
type AccessToken struct {
	Token      string
	ExpireTime string
}

//NewAccessToken ***
func NewAccessToken(token, expireTime string) *AccessToken {
	o := &AccessToken{
		Token:      token,
		ExpireTime: expireTime,
	}
	return o
}

// ToString get the string
func (o *AccessToken) ToString() string {
	return fmt.Sprintf("AccessToken [Token=%s, ExpireTime=%s]", o.Token, o.ExpireTime)
}
