package bimface

import (
	"fmt"
	"testing"

	"github.com/chnykn/bimface/v3/config"
	"github.com/chnykn/bimface/v3/consts"
)

const (
	appKey    = "aaHvGIzGvn5P64JwBEg3L68fIpXqllOp"
	appSecret = "3usnqFp5UBJU30VP0ttziq3hI8qAoQlP"
)

func TestCatalog(t *testing.T) {

	client := NewCatalog(appKey, appSecret,
		config.NewEndpoint(consts.APIHost, consts.FileHost))

	//lst, err := client.IntegrationService.GetList()
	lst, err := client.SourceFileService.GetSupportFile()

	if err != nil {
		fmt.Println("ERR:", err.Error())
	}

	fmt.Println(lst)
}
