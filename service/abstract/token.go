// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abstract

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chnykn/bimface/v3/consts"
	"github.com/chnykn/httpkit"
)

const (
	authTokenURI string = "/oauth2/token"
)

type authToken struct {
	Token      string `json:"token"`
	ExpireTime string `json:"expireTime"`
}

type tokenResp struct {
	Code    string     `json:"code"` // must be "success"
	Message string     `json:"message,omitempty"`
	Data    *authToken `json:"data,omitempty"`
}

//=========================================================

func (o *Service) isValidToken() bool {
	if (o.authToken == nil) || (o.authToken.Token == "") {
		return false
	}

	expire, err := time.ParseInLocation(time.DateTime, o.authToken.ExpireTime, time.Local) //"2006-01-02 15:04:05"
	if err == nil {
		now := time.Now().Add(time.Minute) //过期时间 提早 一分钟
		return now.Before(expire)
	}

	return false
}

func (o *Service) grantToken() (*authToken, error) {

	httpClnt := httpkit.NewClient(func(httpReq *http.Request) error {
		key := o.Credential.AppKey + ":" + o.Credential.AppSecret
		authorization := "Basic" + " " + base64.StdEncoding.EncodeToString([]byte(key))

		httpReq.Header.Set("Authorization", authorization)
		return nil
	})

	//buff := bytes.NewBuffer(nil)
	var buff []byte //也可以直接使用 *[]byte ，必须是指针

	_, err := httpClnt.Post(o.Endpoint.APIHost+authTokenURI, httpkit.NewRespBody(&buff))
	if err != nil {
		return nil, err
	}

	//----------

	resp := new(tokenResp)
	//err = json.Unmarshal(buff.Bytes(), resp)
	err = json.Unmarshal(buff, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != consts.SUCCESS {
		return nil, fmt.Errorf("grantToken ERR: %s", resp.Message)
	}

	return resp.Data, nil
}
