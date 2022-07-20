package integration

import (
	"github.com/chnykn/bimface/v2/bean/request"
	"github.com/chnykn/bimface/v2/bean/response"
)

func (o *Service) GetList() ([]*response.FileIntegrateDetailBean, error) { //appKey string
	qry := &request.IntegrateQueryRequest{
		AppKey:        "",
		ProjectId:     0,
		FileName:      "",
		SourceId:      "",
		IntegrateId:   "",
		IntegrateType: "",
		PageNo:        0,
		PageSize:      300,
		Status:        0,
		SortType:      "",
		StartDate:     "",
		EndDate:       "",
	}

	var err error
	var arr = make([]*response.FileIntegrateDetailBean, 0)
	var res *response.FileIntegrateDetailBeanPageList
	for {
		res, err = o.GetDetails(qry)
		if err == nil {
			arr = append(arr, res.List...)

			page := res.PageList.Page
			if page.PageNo < page.NextPage {
				qry.PageNo = page.PageNo + 1
			} else {
				break
			}
		} else {
			break
		}
	}

	return arr, err
}
