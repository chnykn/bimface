// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bimface

import (
	"github.com/chnykn/bimface/v2/config"
	"github.com/chnykn/bimface/v2/consts"
	"github.com/chnykn/bimface/v2/service"
	"github.com/chnykn/bimface/v2/service/bake"
	"github.com/chnykn/bimface/v2/service/comparison"
	"github.com/chnykn/bimface/v2/service/databag"
	"github.com/chnykn/bimface/v2/service/integration"
	"github.com/chnykn/bimface/v2/service/modelfile"
	"github.com/chnykn/bimface/v2/service/sharelink"
	"github.com/chnykn/bimface/v2/service/sourcefile"
	"github.com/chnykn/bimface/v2/utils"
)

//TODO: 模型信息和构件属性查询、构件空间关系计算、导出数据包相关、烘焙相关、rfa构件数据

// Client for bimface SDK
type Client struct {
	credential    *config.Credential
	endpoint      *config.Endpoint
	serviceClient *utils.ServiceClient

	AccessTokenService *service.AccessTokenService
	ViewTokenService   *service.ViewTokenService

	SourceFileService  *sourcefile.Service
	ModelFileService   *modelfile.Service
	IntegrationService *integration.Service
	ShareLinkService   *sharelink.Service
	ComparisonService  *comparison.Service
	DataBagService     *databag.Service
	BakeService        *bake.Service
}

// NewClient create an bimface client.
func NewClient(appKey string, appSecret string, endpoint *config.Endpoint) *Client {
	if endpoint == nil {
		endpoint = config.NewEndpoint(consts.APIHost, consts.FileHost)
	}

	o := &Client{
		credential:    config.NewCredential(appKey, appSecret),
		endpoint:      endpoint,
		serviceClient: utils.NewServiceClient(),
	}
	o.AccessTokenService = service.NewAccessTokenService(o.serviceClient, o.endpoint, o.credential)
	o.ViewTokenService = service.NewViewTokenService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)

	o.SourceFileService = sourcefile.NewService(o.serviceClient, o.endpoint, o.AccessTokenService)
	o.ModelFileService = modelfile.NewService(o.serviceClient, o.endpoint, o.AccessTokenService)
	o.IntegrationService = integration.NewService(o.serviceClient, o.endpoint, o.AccessTokenService)
	o.ShareLinkService = sharelink.NewService(o.serviceClient, o.endpoint, o.AccessTokenService)
	o.ComparisonService = comparison.NewService(o.serviceClient, o.endpoint, o.AccessTokenService)
	o.DataBagService = databag.NewService(o.serviceClient, o.endpoint, o.AccessTokenService)
	o.BakeService = bake.NewService(o.serviceClient, o.endpoint, o.AccessTokenService)

	return o
}
