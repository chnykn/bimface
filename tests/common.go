package tests

import (
	"github.com/chnykn/bimface"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/consts"
)

const (
	appKey    = "<yourAppKey>"
	appSecret = "<yourAppSecret>"
)

func getClient() *bimface.Client {
	return bimface.NewClient(appKey, appSecret,
		config.NewEndpoint(consts.APIHost, consts.FileHost))
}
