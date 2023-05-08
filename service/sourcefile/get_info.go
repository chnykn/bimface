package sourcefile

import (
	"fmt"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	fileInfoURI string = "/files/%d"
)

func (o *Service) fileInfoURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.FileHost+fileInfoURI, fileId)
}

// 获取文件信息
func (o *Service) GetInfo(fileId int64) (*response.FileBean, error) {
	result := new(response.FileBean)
	err := o.GET(o.fileInfoURL(fileId), result)

	return result, err
}
