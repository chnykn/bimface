// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"fmt"

	"github.com/imroc/req"

	"github.com/chnykn/bimface/bean/response"
	"github.com/chnykn/bimface/utils"
)

const (
	//PUT GET https://api.bimface.com/files/{fileId}/split
	splitDrawing string = "/files/%d/split"

	//GET https://api.bimface.com/data/v2/files/{fileId}/frames
	splitDrawingFrames string = "/data/v2/files/%d/frames"
)

//---------------------------------------------------------------------

func (o *Service) splitDrawingURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+splitDrawing, fileId)
}

func (o *Service) splitDrawingFramesURL(fileId int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+splitDrawingFrames, fileId)
}

//-----------------------------------------------------------------------------------

func (o *Service) doSplitDrawing(fileId int64, isGet bool) (*response.SplitDrawingStatus, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	var resp *req.Resp

	if isGet {
		resp = o.ServiceClient.Get(o.splitDrawingURL(fileId), headers.Header)
	} else {
		resp = o.ServiceClient.Put(o.splitDrawingURL(fileId), headers.Header)
	}

	result := response.NewSplitDrawingStatus()
	err = utils.RespToBean(resp, result)

	return result, err
}

//发起图纸拆分
func (o *Service) SplitDrawing(fileId int64) (*response.SplitDrawingStatus, error) {
	return o.doSplitDrawing(fileId, false)
}

//获取图纸拆分状态
func (o *Service) GetSplitDrawingStatus(fileId int64) (*response.SplitDrawingStatus, error) {
	return o.doSplitDrawing(fileId, true)
}

//-----------------------------------------------------------------------------------

//获取图纸拆分结果
func (o *Service) GetSplitDrawingFrames(fileId int64) ([]*response.SplitDrawingFrames, error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := utils.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.splitDrawingFramesURL(fileId), headers.Header)

	var result []*response.SplitDrawingFrames
	err = utils.RespToBean(resp, result)

	return result, err
}
