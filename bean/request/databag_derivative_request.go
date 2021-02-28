// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

type DataBagDerivativeRequest struct {
	Config map[string]string `json:"config"`
}

func NewDataBagDerivativeRequest() *DataBagDerivativeRequest {
	o := &DataBagDerivativeRequest{
		Config: make(map[string]string, 0),
	}
	return o
}
