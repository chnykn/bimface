// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bean

import "fmt"

const (
	//CodeSuccess ***
	CodeSuccess = "success"
	//DateFormat ***
	DateFormat = "yyyy-MM-dd HH:mm:ss" //2006-01-02 15:04:05
)

//ResponseResult ***
type ResponseResult struct {
	Code    string `json:"code"` // = "success"
	Message string `json:"message"`
	//Data    interface{} //Data is bean.response.Xxx
}

// NewResponseResult ***
func NewResponseResult(data interface{}) *ResponseResult {
	o := &ResponseResult{
		Code:    CodeSuccess,
		Message: "",
		//Data:    data,
	}

	return o
}

// ToString get the string
func (o *ResponseResult) ToString() string {
	return fmt.Sprintf("ResponseResult [code=%s, message=%s]", o.Code, o.Message)
}
