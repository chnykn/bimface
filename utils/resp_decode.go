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

//RespToResponseResult ***
func RespToResponseResult(resp *req.Resp) (*bean.ResponseResult, error) {
	if resp == nil {
		return nil, fmt.Errorf("RespToResponseResult Error, resp == nil")
	}

	result := new(bean.ResponseResult)
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
	result, err := RespToResponseResult(resp)
	if err != nil {
		return err
	}

	if result.Code != bean.CodeSuccess {
		return fmt.Errorf("RespToBean Error, code=%s, message=%s", result.Code, result.Message)
	}

	rjson, _ := resp.ToString()
	n := strings.Index(rjson, `"data":`)
	if n < 0 {
		return fmt.Errorf(`RespToBean Error, not found "data" in json`)
	}

	rjson = rjson[n+7 : strings.LastIndex(rjson, "}")]
	fmt.Println(rjson)

	return json.Unmarshal([]byte(rjson), returnObj)
}
