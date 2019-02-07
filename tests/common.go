package tests

import (
	"github.com/chnykn/bimface"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/consts"
)

func getClient() *bimface.Client {
	return bimface.NewClient(<yourAppKey>, <yourAppSecret>,
		config.NewEndpoint(consts.APIHost, consts.FileHost))
}
