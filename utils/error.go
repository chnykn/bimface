package utils

// Error define
type Error struct {
	Code    string
	Message string
}

//NewError ***
func NewError(code string, message string) *Error {
	o := &Error{
		Code:    code,
		Message: message,
	}
	return o
}
