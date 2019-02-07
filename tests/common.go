package tests

import (
	"github.com/chnykn/bimface"
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/consts"
)

func getClient() *bimface.Client {
	return bimface.NewClient("purREs2GDIdCL81tUVSS6pAMfkZ59KNi", "954ZyagB5vskJXKb2tZWSRnhJhRNKQ96",
		config.NewEndpoint(consts.APIHost, consts.FileHost))
}
