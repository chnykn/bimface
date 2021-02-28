// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package request

import "bytes"

type FileUploadRequest struct {
	Name     string `json:"name"`
	URL      string `json:"url,omitempty"`
	Buffer   bytes.Buffer
	SourceId string `json:"sourceId,omitempty"`
}

//IsByURL ***
func (o *FileUploadRequest) IsByURL() bool {
	return (o.URL != "")
}

func NewFileUploadRequest() *FileUploadRequest {
	o := &FileUploadRequest{}
	return o
}
