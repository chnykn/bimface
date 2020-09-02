package tests

import (
	"github.com/chnykn/bimface"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/consts"
)

const (
	appKey    = "ipurREs2GDIdCL81tUVSS6pAMfkZ59KN"
	appSecret = "6954ZyagB5vskJXKb2tZWSRnhJhRNKQ9"
)

func getClient() *bimface.Client {
	return bimface.NewClient(appKey, appSecret,
		config.NewEndpoint(consts.APIHost, consts.FileHost))
}
