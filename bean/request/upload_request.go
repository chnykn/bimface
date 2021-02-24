// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package request

import "bytes"

//UploadRequest ***
type UploadRequest struct {
	Name     string `json:"name"`
	SourceId string `json:"sourceId,omitempty"`
	URL      string `json:"url,omitempt"`
	Buffer   bytes.Buffer
}

//IsByURL ***
func (o *UploadRequest) IsByURL() bool {
	return (o.URL != "")
}
