// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

//TranslateSource ***
type TranslateSource struct {
	FileID     int64  `json:"fileId"`               //必填
	Compressed bool   `json:"compressed,omitempty"` //是否为压缩文件 默认为false
	RootName   string `json:"rootName,omitempty"`   //如果是压缩文件，必须指定压缩包中哪一个是主文件
}

//NewTranslateSource ***
func NewTranslateSource(fileID int64) *TranslateSource {
	o := &TranslateSource{
		FileID:     fileID,
		Compressed: false,
		RootName:   "",
	}
	return o
}

//---------------------------------------------------------------------

const (
	defualtTranslatePriority int = 2
)

//TranslateRequest ***
type TranslateRequest struct {
	Source   *TranslateSource `json:"source"`
	Priority int              `json:"priority"` //[1,2,3] 数字越大，优先级越低
	Callback string           `json:"callback,omitempty"`
}

//NewTranslateRequest ***
func NewTranslateRequest(sourceFileID int64) *TranslateRequest {
	o := &TranslateRequest{
		Source:   NewTranslateSource(sourceFileID),
		Priority: defualtTranslatePriority,
		Callback: "",
	}
	return o
}
