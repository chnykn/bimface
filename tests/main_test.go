package tests

import (
	"fmt"
	"testing"

	"github.com/chnykn/bimface/v3"
	"github.com/chnykn/bimface/v3/config"
	"github.com/chnykn/bimface/v3/consts"
)

func TestCatalog(t *testing.T) {
	//fill in your own appKey, appSecret
	client := bimface.NewCatalog(appKey, appSecret,
		config.NewEndpoint(consts.APIHost, consts.FileHost))

	//lst, err := client.IntegrationService.GetList()
	lst, err := client.SourceFileService.GetSupportFile()

	if err != nil {
		fmt.Println("ERR:", err.Error())
	}

	fmt.Println(lst)
}
