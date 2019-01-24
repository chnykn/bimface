package response

import "fmt"

//ShareLink ***
type ShareLink struct {
	URL        string
	ExpireTime string
}

//NewShareLink ***
func NewShareLink(url, expireTime string) *ShareLink {
	o := &ShareLink{
		URL:        url,
		ExpireTime: expireTime,
	}
	return o
}

// ToString get the string
func (o *ShareLink) ToString() string {
	return fmt.Sprintf("ShareLink [URL=%s, ExpireTime=%s]", o.URL, o.ExpireTime)
}
