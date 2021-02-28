// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

const defaultUserAgent = ""

/**
//Config ***
type Config struct {
	UserAgent           string // defaultUserAgent
	MaxIdleConnections  int64  // 32
	KeepAliveDurationNs int64  // 300000
	MaxRequests         int64  // 64
	MaxRequestsPerHost  int64  // 5
	ConnectTimeout      int64  // 10
	ReadTimeout         int64  // 0
	WriteTimeout        int64  // 30
}

//NewConfig ***
func NewConfig() *Config {
	o := &Config{
		UserAgent:           defaultUserAgent,
		MaxIdleConnections:  32,
		KeepAliveDurationNs: 300000,
		MaxRequests:         64,
		MaxRequestsPerHost:  5,
		ConnectTimeout:      10,
		ReadTimeout:         0,
		WriteTimeout:        30,
	}
	return o
}
**/
