package tests

import (
	"fmt"
	"github.com/chnykn/bimface/v3/bean/request"
	"testing"

	"github.com/chnykn/bimface/v3"
	"github.com/chnykn/bimface/v3/config"
	"github.com/chnykn/bimface/v3/consts"
)

func getElemTree(catalog *bimface.Catalog) {
	elemTree, err := catalog.IntegrationService.GetElementTree(integrateId, "specialty", nil, nil)
	if err != nil {
		fmt.Println("ERR:", err.Error())
		return
	}

	saveToJsonFile(elemTree, "e:/intgr-spec-tree.json")
	fmt.Println(elemTree)
}

func queryEleme(catalog *bimface.Catalog) {
	query := &request.DSLCondition{
		Match: map[string]any{"familyType": "JZ-砌块墙-300mm"},
		//Contain: map[string]any{"floor": "B01", "familyType": "标准"},
		//BoolAnd: []request.DSLCondition{
		//	{Match: map[string]any{"categoryId": "id111"}},
		//	{Match: map[string]any{"boundingBox.min.x": 167899.9999999998}},
		//},
		//BoolOr: []request.DSLCondition{
		//	{Match: map[string]any{"productID": "KDKE-B-9947-#kL5"}},
		//	{
		//		BoolAnd: []request.DSLCondition{
		//			{Match: map[string]any{"productID": "JODL-X-1937-#pV7"}},
		//			{Match: map[string]any{"price": 30}},
		//		},
		//	},
		//},
	}
	elemIds, err := catalog.IntegrationService.QueryElementIds([]int64{integrateId}, query, false)
	if err != nil {
		fmt.Println("ERR:", err.Error())
		return
	}

	fmt.Println(elemIds)
}

func TestGetIntgrTree(t *testing.T) {
	//fill in your own appKey, appSecret
	catalog := bimface.NewCatalog(appKey, appSecret,
		config.NewEndpoint(consts.APIHost, consts.FileHost))

	//getElemTree(catalog)
	queryEleme(catalog)
}
