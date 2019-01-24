package bean

import "fmt"

const (
	//CodeSuccess ***
	CodeSuccess = "success"
	//DateFormat ***
	DateFormat = "yyyy-MM-dd HH:mm:ss" //2006-01-02 15:04:05
)

//GeneralResponse ***
type GeneralResponse struct {
	Code    string // = "success"
	Message string
	Data    interface{} //Data is bean.response.Xxx
}

// NewGeneralResponse ***
func NewGeneralResponse(data interface{}) *GeneralResponse {
	o := &GeneralResponse{
		Code:    CodeSuccess,
		Message: "",
		Data:    data,
	}

	return o
}

// ToString get the string
func (o *GeneralResponse) ToString() string {
	return fmt.Sprintf("GeneralResponse [code=%s, message=%s, data=%v]", o.Code, o.Message, o.Data)
}
