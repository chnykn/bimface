// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

type ElementPropertyValueOverride struct {
	ValueToMatch    string `json:"valueToMatch"`
	ValueToOverride string `json:"valueToOverride"`
}

type ElementPropertyOverride struct {
	KeyToMatch    string `json:"keyToMatch"`
	KeyToOverride string `json:"keyToOverride"`

	TargetFileIds  []interface{}           `json:"targetFileIds"`
	ValueOverrides []PropertyValueOverride `json:"valueOverrides"` //
}

//--------------------------------

type PropertyValueOverride ElementPropertyValueOverride

type PropertyOverride ElementPropertyOverride
