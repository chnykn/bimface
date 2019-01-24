package utils

import (
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
func CheckFileName(fileName string) *Error {
	if fileName == "" {
		return NewError("invalid.filename", "File name must not be empty.")
	}

	if len(fileName) > 256 {
		return NewError("invalid.filename", "File name too long, no more than 256 characters.")
	}

	suffix := getSuffix(fileName)
	if suffix == "" {
		return NewError("invalid.filename", "File name has no suffix.")
	}

	if containsIllegalChar(fileName) {
		return NewError("invalid.filename", "File name contains illegal character.")
	}

	return nil
}

//CheckFileType ***
func CheckFileType(allSupportedType []string, fileName string) *Error {
	suffix := strings.ToLower(getSuffix(fileName))

	for _, typ := range allSupportedType {
		if suffix == strings.ToLower(typ) {
			return NewError("invalid.filetype", "File type not supported.")
		}
	}

	return nil
}
