package tests

import (
	"fmt"
	"github.com/chnykn/bimface/v3"
	"github.com/chnykn/bimface/v3/config"
	"github.com/chnykn/bimface/v3/consts"
	"testing"
)

func testElemTree(catalog *bimface.Catalog) {
	elemTree, err := catalog.IntegrationService.GetElementTree(integrateId, "specialty", nil, nil)

	if err != nil {
		fmt.Println("ERR:", err.Error())
	}

	saveToJsonFile(elemTree, "e:/intgr-spec-tree.json")
	fmt.Println(elemTree)
}

func TestGetIntgrTree(t *testing.T) {
	//fill in your own appKey, appSecret
	catalog := bimface.NewCatalog(appKey, appSecret,
		config.NewEndpoint(consts.APIHost, consts.FileHost))

	testElemTree(catalog)
}
