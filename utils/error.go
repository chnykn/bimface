// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

// Error define
type Error struct {
	Code    string
	Message string
}

//NewError ***
func NewError(code string, message string) *Error {
	o := &Error{
		Code:    code,
		Message: message,
	}
	return o
}
