// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bean

import (
	"encoding/json"
	"fmt"

	"github.com/chnykn/bimface/v3/consts"
)

// RespResult ***
type RespResult struct {
	Code    string          `json:"code"` // = "success"
	Message string          `json:"message,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"` //Data is bean.response.Xxx
}

// NewRespResult ***
func NewRespResult() *RespResult {
	o := &RespResult{
		Code:    consts.SUCCESS,
		Message: "",
		Data:    nil,
	}

	return o
}

// ToString get the string
func (o *RespResult) ToString() string {
	return fmt.Sprintf("RespResult [code=%s, message=%s]", o.Code, o.Message)
}
