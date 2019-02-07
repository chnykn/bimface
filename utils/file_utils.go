// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"path"
	"strings"
)

var illegalChars = [10]string{"/", "\n", "*", "\\", "<", ">", "|", "\"", ":", "?"}

func containsIllegalChar(value string) bool {
	for _, c := range illegalChars {
		n := strings.Count(value, c)
		if n > 0 {
			return true
		}
	}

	return false
}

func getSuffix(fileName string) string {
	var suffix string

	suffix = path.Ext(fileName)
	suffix = strings.TrimLeft(suffix, ".")

	return suffix
}

//CheckFileName ***
func CheckFileName(fileName string) error {
	if fileName == "" {
		return fmt.Errorf("file name must not be empty")
	}

	if len(fileName) > 256 {
		return fmt.Errorf("file name too long, no more than 256 characters")
	}

	suffix := getSuffix(fileName)
	if suffix == "" {
		return fmt.Errorf("file name has no suffix")
	}

	if containsIllegalChar(fileName) {
		return fmt.Errorf("file name contains illegal character")
	}

	return nil
}

//CheckFileType ***
func CheckFileType(allSupportedType []string, fileName string) error {
	suffix := strings.ToLower(getSuffix(fileName))

	for _, typ := range allSupportedType {
		if suffix == strings.ToLower(typ) {
			return fmt.Errorf("file type not supported")
		}
	}

	return nil
}
