package bimface

import (
	"fmt"
	"github.com/chnykn/bimface/v2/config"
	"github.com/chnykn/bimface/v2/consts"
	"testing"
)

func TestBimFaceClient(t *testing.T) {

	client := NewClient("xxxxxxx",
		"xxxxxx",
		config.NewEndpoint(consts.APIHost, consts.FileHost))

	lst, err := client.IntegrationService.GetList()

	if err != nil {
		fmt.Println("ERR:", err.Error())
	}

	fmt.Println(lst)
}
