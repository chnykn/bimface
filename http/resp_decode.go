// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/chnykn/bimface/bean"
	"github.com/chnykn/bimface/utils"
	"reflect"

	"github.com/goinggo/mapstructure"
	"github.com/imroc/req"
)

//errorInfo ***
type errorInfo struct {
	Timestamp int64
	Status    int
	Error     string
	Message   string
}

//RespToGeneralResponse ***
func RespToGeneralResponse(resp *req.Resp) (*bean.GeneralResponse, *utils.Error) {
	if resp == nil {
		return nil, utils.NewError("failed", "resp == nil @ RespToGeneralResponse")
	}

	result := new(bean.GeneralResponse)
	err := resp.ToJSON(result)
	if err != nil {
		return nil, utils.NewError("failed", "resp.ToJSON(result) @ RespToGeneralResponse")
	}

	if result.Code == "" {
		errInfo := new(errorInfo)
		resp.ToJSON(errInfo)
		result.Code = errInfo.Error
	}
	return result, nil
}

//------------------------------------------------------------------------

func decodeSlice(data interface{}, val reflect.Value) interface{} {

	valType := reflect.Indirect(val).Type()
	valSliceType := reflect.SliceOf(valType)
	valSlice := reflect.MakeSlice(valSliceType, 0, 0) // Make a slice to hold result

	dataVal := reflect.Indirect(reflect.ValueOf(data))

	for i := 0; i < dataVal.Len(); i++ {
		currData := dataVal.Index(i).Interface()
		currItem := reflect.New(valType).Interface() // *result

		err := mapstructure.Decode(currData, currItem)
		if err == nil {
			valSlice = reflect.Append(valSlice, reflect.ValueOf(currItem).Elem())
		}
	}

	return valSlice.Interface()
}

//RespToBeans : the obj must be pointer for reflect
func RespToBeans(resp *req.Resp, obj interface{}) (interface{}, *utils.Error) {
	gresp, err := RespToGeneralResponse(resp)
	if err != nil {
		return nil, err
	}

	if gresp.Code == bean.CodeSuccess {
		result := decodeSlice(gresp.Data, reflect.ValueOf(obj))
		return result, nil
	}

	return nil, utils.NewError(gresp.Code, gresp.Message)
}

//-----------------------------------------------------------------------

//RespToBean : the obj must be pointer, and will be populated by resp
func RespToBean(resp *req.Resp, obj interface{}) *utils.Error {
	gresp, err := RespToGeneralResponse(resp)
	if err != nil {
		return err
	}

	if gresp.Code == bean.CodeSuccess {
		err := mapstructure.Decode(gresp.Data, obj)
		if err != nil {
			return utils.NewError(err.Error(), "mapstructure.Decode(gresp.Data, obj) @ RespToBean")
		}
		return nil
	}

	return utils.NewError(gresp.Code, gresp.Message)
}
