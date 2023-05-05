// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bimface

import (
	"github.com/chnykn/bimface/v3/config"
	"github.com/chnykn/bimface/v3/consts"
	"github.com/chnykn/bimface/v3/service/abstract"
	"github.com/chnykn/bimface/v3/service/bake"
	"github.com/chnykn/bimface/v3/service/comparison"
	"github.com/chnykn/bimface/v3/service/databag"
	"github.com/chnykn/bimface/v3/service/integration"
	"github.com/chnykn/bimface/v3/service/modelfile"
	"github.com/chnykn/bimface/v3/service/sharelink"
	"github.com/chnykn/bimface/v3/service/sourcefile"
	"github.com/chnykn/bimface/v3/service/viewtoken"
)

//http://static.bimface.com/restful-apidoc/dist/index.html

// Catalog for bimface SDK
type Catalog struct {
	credential *config.Credential
	endpoint   *config.Endpoint
	absService *abstract.Service

	ViewTokenService   *viewtoken.Service
	SourceFileService  *sourcefile.Service
	ModelFileService   *modelfile.Service
	IntegrationService *integration.Service
	ShareLinkService   *sharelink.Service
	ComparisonService  *comparison.Service
	DataBagService     *databag.Service
	BakeService        *bake.Service
}

// NewCatalog create an bimface client.
func NewCatalog(appKey string, appSecret string, endpoint *config.Endpoint) *Catalog {

	credential := config.NewCredential(appKey, appSecret)

	if endpoint == nil {
		endpoint = config.NewEndpoint(consts.APIHost, consts.FileHost)
	}

	o := &Catalog{
		credential: credential,
		endpoint:   endpoint,
		absService: abstract.NewService(credential, endpoint),
	}

	o.ViewTokenService = viewtoken.NewService(o.absService)
	o.SourceFileService = sourcefile.NewService(o.absService)
	o.ModelFileService = modelfile.NewService(o.absService)
	o.IntegrationService = integration.NewService(o.absService)
	o.ShareLinkService = sharelink.NewService(o.absService)
	o.ComparisonService = comparison.NewService(o.absService)
	o.DataBagService = databag.NewService(o.absService)
	o.BakeService = bake.NewService(o.absService)

	return o
}
