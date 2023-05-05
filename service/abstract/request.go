package abstract

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chnykn/bimface/v3/bean"
	"github.com/chnykn/bimface/v3/consts"
	"github.com/chnykn/httpkit"
)

func (o *Service) beforeReq(httpReq *http.Request) error {
	if !o.isValidToken() {
		token, err := o.grantToken()
		if err != nil {
			return err
		}
		o.authToken = token
	}

	httpReq.Header.Set("Authorization", "Bearer "+o.authToken.Token)
	return nil
}

func (o *Service) request(url string, method string, result any, v ...any) error {
	buff := bytes.NewBuffer(nil)
	v = append(v, httpkit.NewRespBody(buff))

	_, err := o.httpClint.Request(url, method, v...)
	if err != nil {
		return err
	}

	respRes := bean.NewRespResult()
	err = json.Unmarshal(buff.Bytes(), respRes)
	if err != nil {
		return err
	}

	if respRes.Code != consts.SUCCESS {
		return fmt.Errorf("AbstractService.request Error, code=%s, message=%s", respRes.Code, respRes.Message)
	}

	if result != nil {
		if len(respRes.Data) > 0 {
			err = json.Unmarshal(respRes.Data, result)
		} else {
			return fmt.Errorf("AbstractService.request Error, data is null")
		}
	}

	return err
}

// GET ***
func (o *Service) GET(url string, result any, v ...any) error {
	return o.request(url, httpkit.MethodGet, result, v...)
}

// POST ***
func (o *Service) POST(url string, result any, v ...any) error {
	return o.request(url, httpkit.MethodPost, result, v...)
}

// PUT ***
func (o *Service) PUT(url string, result any, v ...any) error {
	return o.request(url, httpkit.MethodPut, result, v...)
}

// DELETE ***
func (o *Service) DELETE(url string, result any, v ...any) error {
	return o.request(url, httpkit.MethodDelete, result, v...)
}
