package comm

import (
	"fmt"
	"strconv"

	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/bimface/v3/service/abstract"
	"github.com/chnykn/httpkit"
)

const (
	// 查询符合条件的构件ID列表
	// https://bimface.com/docs/model-data-service/v1/api-reference/getElementsUsingPOST_2.html
	// 查询符合条件的构件ID列表，支持基于DSL进行组合查询，具体可参考[文档](https://bimface.com/docs/model-data-service/v1/developers-guide/query-instruction.html)。
	//  若需根据修改后的属性值查询，需要设置请求参数includeOverrides的值为true。
	queryElementIdsURI string = "/data/v2/query/elementIds"
)

func queryElementIdsURL(service *abstract.Service, includeOverrides bool) string {
	result := fmt.Sprintf(service.Endpoint.APIHost + queryElementIdsURI)
	if includeOverrides {
		result = result + "?includeOverrides=true"
	}
	return result
}

// https://bimface.com/docs/model-data-service/v1/developers-guide/query-instruction.html
/*
字段			类型			必填		描述
targetType	String		Y		查询目标类型，只能是file或integration
targetIds	String[]	Y		查询目标ID列表
query		Object		Y		查询条件实体，由match、contain、boolAnd、boolOr组成
includeOverrides				是否根据修改后的属性值查询，默认为false	boolean
*/

func QueryElementIds(service *abstract.Service, targetType string, targetIds []int64,
	query *request.DSLCondition, includeOverrides bool) (response.ElementIdsArrayBean, error) {

	targetStrIds := make([]string, len(targetIds))
	for i := 0; i < len(targetIds); i++ {
		targetStrIds[i] = strconv.FormatInt(targetIds[i], 10)
	}

	dsl := &request.QueryDSLRequest{
		TargetType: targetType,
		TargetIds:  targetStrIds,
		Query:      query,
	}

	//----------------

	result := make(response.ElementIdsArrayBean, 0)

	body := httpkit.JsonReqBody(dsl)
	err := service.POST(queryElementIdsURL(service, includeOverrides), &result, body)

	return result, err
}
