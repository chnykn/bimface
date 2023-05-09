package modelfile

import (
	"strconv"
	"strings"

	"github.com/chnykn/bimface/v3/bean/response"
)

const (
	// 查询指定构件属性的所有值，目前目标类型仅支持file
	// https://bimface.com/docs/model-data-service/v1/api-reference/getPropertyValuesUsingGET.html
	queryPropertyValueURI string = "/data/v2/query/propertyValues"
)

func (o *Service) queryPropertyValuesURL(fileIds []int64, properties []string, includeOverrides bool) string {
	result := queryPropertyValueURI
	result = result + "?properties=" + strings.Join(properties, ",")

	targetIds := make([]string, len(fileIds))
	for i := range fileIds {
		targetIds[i] = strconv.FormatInt(fileIds[i], 10)
	}
	result = result + "&targetIds=" + strings.Join(targetIds, ",")

	if includeOverrides {
		result = result + "&includeOverrides=true"
	}
	return result
}

func (o *Service) QueryPropertyValues(fileIds []int64, properties []string,
	includeOverrides bool) (*response.PropertyValuesBeanList, error) {

	result := new(response.PropertyValuesBeanList)
	err := o.GET(o.queryPropertyValuesURL(fileIds, properties, includeOverrides), result)

	return result, err
}
