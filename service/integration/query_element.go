package integration

import (
	"github.com/chnykn/bimface/v3/bean/request"
	"github.com/chnykn/bimface/v3/bean/response"
	"github.com/chnykn/bimface/v3/service/comm"
)

func (o *Service) QueryElementIds(integrateIds []int64, query *request.DSLCondition,
	includeOverrides bool) (*response.ElementIdsArrayBean, error) {

	return comm.QueryElementIds(o.Service, "integration", integrateIds, query, includeOverrides)
}
