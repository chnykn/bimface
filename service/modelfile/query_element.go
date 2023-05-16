package modelfile

import (
	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/bimface/v3/service/comm"
)

/*
const (
	// 查询符合条件的构件ID列表
	// https://bimface.com/docs/model-data-service/v1/api-reference/getElementsUsingPOST_2.html
	// 查询符合条件的构件ID列表，支持基于DSL进行组合查询，具体可参考[文档](https://bimface.com/docs/model-data-service/v1/developers-guide/query-instruction.html)。
	//  若需根据修改后的属性值查询，需要设置请求参数includeOverrides的值为true。
	queryElementIdsURI string = "/data/v2/query/elementIds"
)

func (o *Service) queryElementIdsURL(includeOverrides bool) string {
	result := fmt.Sprintf(o.Endpoint.APIHost + queryElementIdsURI)
	if includeOverrides {
		result = result + "?includeOverrides=true"
	}
	return result
}

func (o *Service) QueryElementIds(query *request.DSLCondition, fileIds []int64, includeOverrides bool) (*response.ElementIdsBean, error) {

	targetIds := make([]string, len(fileIds))
	for i := 0; i < len(fileIds); i++ {
		targetIds[i] = strconv.FormatInt(fileIds[i], 10)
	}

	dsl := &request.QueryDSLRequest{
		TargetType: "file",
		TargetIds:  targetIds,
		Query:      query,
	}

	//----------------

	result := new(response.ElementIdsBean)

	body := httpkit.JsonReqBody(dsl)
	err := o.POST(o.queryElementIdsURL(includeOverrides), &result, body)

	return result, err
}
*/

func (o *Service) QueryElementIds(fileIds []int64, query *request.DSLCondition, includeOverrides bool) (*response.ElementIdsBean, error) {
	return comm.QueryElementIds(o.Service, "file", fileIds, query, includeOverrides)
}
