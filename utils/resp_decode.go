// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/chnykn/bimface/bean"

	"github.com/imroc/req"
)

//errorInfo ***
type errorInfo struct {
	Timestamp int64
	Status    int
	Error     string
	Message   string
}

//RespToResult ***
func RespToResult(resp *req.Resp) (*bean.RespResult, error) {
	if resp == nil {
		return nil, fmt.Errorf("RespToResult Error, resp == nil")
	}

	result := new(bean.RespResult)
	err := resp.ToJSON(result)
	if err != nil {
		return nil, err
	}

	if result.Code == "" {
		errInfo := new(errorInfo)
		resp.ToJSON(errInfo)
		result.Code = errInfo.Error
	}

	return result, nil
}

//RespToBean : the obj must be pointer, and will be populated by resp
func RespToBean(resp *req.Resp, returnObj interface{}) error {
	result, err := RespToResult(resp)
	if err != nil {
		return err
	}

	if result.Code != bean.CodeSuccess {
		return fmt.Errorf("RespToBean Error, code=%s, message=%s", result.Code, result.Message)
	}

	text, _ := resp.ToString()
	n := strings.Index(text, `"data":`)
	if n < 0 {
		return fmt.Errorf(`RespToBean Error, not found "data" in json`)
	}

	text = text[n+7 : strings.LastIndex(text, "}")]
	//fmt.Println(text)

	return json.Unmarshal([]byte(text), returnObj)
}
