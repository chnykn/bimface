// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

//CompareSource ***
type CompareSource struct {
	PreviousFileId  int64 `json:"previousFileId"`  //变更前文件Id，如果为新增文件，则为null
	FollowingFileId int64 `json:"followingFileId"` //变更后文件Id，如果为删除文件，则为null
}

//NewCompareSource ***
func NewCompareSource(previousFileId, followingFileId int64) *CompareSource {
	o := &CompareSource{
		PreviousFileId:  previousFileId,
		FollowingFileId: followingFileId,
	}
	return o
}

//---------------------------------------------------------------------

const (
	defualtComparePriority int = 2
)

//CompareRequest ***
type CompareRequest struct {
	Name     string           `json:"Name"`     //用户指定对比后的模型的名字
	Priority int              `json:"priority"` //[1,2,3] 数字越大，优先级越低
	Callback string           `json:"callback,omitempty"`
	SourceId string           `json:"sourceId,omitempty"` //第三方应用自己的Id
	Sources  []*CompareSource `json:"sources"`
}

//NewCompareRequest ***
func NewCompareRequest() *CompareRequest {
	o := &CompareRequest{
		Priority: defualtComparePriority,
		Sources:  make([]*CompareSource, 0),
	}
	return o
}
