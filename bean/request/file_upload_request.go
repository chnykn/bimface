package request

import (
	"bimface/http"
	"os"
	"path"
)

//FileUploadRequest ***
type FileUploadRequest struct {
	Name          string
	SourceID      string
	URL           string
	ContentLength int64  // multipart.FileHeader.Size
	InputStream   []byte // *multipart.FileHeader
	Bucket        string
	ObjectKey     string
}

//NewFileUploadRequest ***
func NewFileUploadRequest(localFile string) *FileUploadRequest {

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

	o := &FileUploadRequest{
		Name:          http.EncodeURI(fileName),
		ContentLength: fileLen,
		InputStream:   fileBuf,
	}
	return o
}

//---------------------------------------------------------------------

//IsByURL ***
func (o *FileUploadRequest) IsByURL() bool {
	return (o.URL != "")
}

//IsByOSS ***
func (o *FileUploadRequest) IsByOSS() bool {
	return (o.Bucket != "") && (o.ObjectKey != "")
}
