// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bimface

import (
	"github.com/chnykn/bimface/config"
	"github.com/chnykn/bimface/consts"
	"github.com/chnykn/bimface/service"
	"github.com/chnykn/bimface/utils"
)

// Client for binface SDK
type Client struct {
	credential *config.Credential
	endpoint   *config.Endpoint

	serviceClient      *utils.ServiceClient
	AccessTokenService *service.AccessTokenService
	SupportFileService *service.SupportFileService

	AppendFileService    *service.AppendFileService
	CategoryTreeService  *service.CategoryTreeService
	CompareService       *service.CompareService
	DownloadService      *service.DownloadService
	DrawingSheetsService *service.DrawingSheetsService
	ElementService       *service.ElementService
	FloorService         *service.FloorService
	IntegrateService     *service.IntegrateService
	IntgrTreeService     *service.IntgrTreeService
	DatabagService       *service.DatabagService
	PropertyService      *service.PropertyService
	ShareLinkService     *service.ShareLinkService
	TranslateService     *service.TranslateService
	UploadService        *service.UploadService
	ViewTokenService     *service.ViewTokenService
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
	o.SupportFileService = service.NewSupportFileService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)

	o.AppendFileService = service.NewAppendFileService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService, o.SupportFileService)
	o.CategoryTreeService = service.NewCategoryTreeService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.CompareService = service.NewCompareService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.DownloadService = service.NewDownloadService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.ElementService = service.NewElementService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.DrawingSheetsService = service.NewDrawingSheetsService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.FloorService = service.NewFloorService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.IntegrateService = service.NewIntegrateService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.IntgrTreeService = service.NewIntgrTreeService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.DatabagService = service.NewDatabagService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.PropertyService = service.NewPropertyService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.ShareLinkService = service.NewShareLinkService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.TranslateService = service.NewTranslateService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.UploadService = service.NewUploadService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)
	o.ViewTokenService = service.NewViewTokenService(o.serviceClient, o.endpoint, o.credential, o.AccessTokenService)

	return o
}
