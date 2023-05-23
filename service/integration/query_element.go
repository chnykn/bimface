package integration

import (
	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/bimface/v3/service/comm"
)

/*
字段				类型			必填		描述
integrateIds	int64[]		Y		查询目标ID列表
query			Object		Y		查询条件实体，由match、contain、boolAnd、boolOr组成
includeOverrides					是否根据修改后的属性值查询，默认为false	boolean
*/

func (o *Service) QueryElementIds(integrateIds []int64, query *request.DSLCondition,
	includeOverrides bool) (response.ElementIdsArrayBean, error) {

	return comm.QueryElementIds(o.Service, "integration", integrateIds, query, includeOverrides)
}
