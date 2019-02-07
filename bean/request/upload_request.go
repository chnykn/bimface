// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

import (
	"os"
	"path"

	"github.com/chnykn/bimface/utils"
)

//UploadRequest ***
type UploadRequest struct {
	Name          string `json:"name"`
	SourceID      string `json:"sourceId,omitempty"`
	URL           string `json:"url,omitempt"`
	ContentLength int64  `json:"contentLength,omitempt"` // multipart.FileHeader.Size
	InputStream   []byte `json:"-"`                      // *multipart.FileHeader
	Bucket        string `json:"bucket,omitempty"`
	ObjectKey     string `json:"objectKey,omitempty"`
}

//NewUploadRequest ***
func NewUploadRequest(localFile string) *UploadRequest {

	var fileLen int64
	var fileBuf []byte
	var fileName string

	file, err := os.Open(localFile)
	if err == nil {
		fileInfo, err := file.Stat()
		if err == nil {
			fileLen = fileInfo.Size()
			if fileLen > 0 {
				fileBuf = make([]byte, fileLen)
				file.Read(fileBuf)

				fileName = path.Base(localFile)
			}
		}
	}

	if fileLen <= 0 {
		fileBuf = nil
	}

	o := &UploadRequest{
		Name:          utils.EncodeURI(fileName),
		ContentLength: fileLen,
		InputStream:   fileBuf,
	}
	return o
}

//---------------------------------------------------------------------

//IsByURL ***
func (o *UploadRequest) IsByURL() bool {
	return (o.URL != "")
}

//IsByOSS ***
func (o *UploadRequest) IsByOSS() bool {
	return (o.Bucket != "") && (o.ObjectKey != "")
}
