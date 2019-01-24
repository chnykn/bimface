// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

//CompareSource ***
type CompareSource struct {
	PreviousFileID  int64 //变更前文件ID，如果为新增文件，则为null
	FollowingFileID int64 //变更后文件ID，如果为删除文件，则为null
}

//NewCompareSource ***
func NewCompareSource(previousFileID, followingFileID int64) *CompareSource {
	o := &CompareSource{
		PreviousFileID:  previousFileID,
		FollowingFileID: followingFileID,
	}
	return o
}

//---------------------------------------------------------------------

const (
	defualtComparePriority int = 2
)

//CompareRequest ***
type CompareRequest struct {
	Name     string //用户指定对比后的模型的名字
	Priority int    //[1,2,3] 数字越大，优先级越低
	Callback string
	SourceID string //第三方应用自己的ID
	Sources  []*CompareSource
}

//NewCompareRequest ***
func NewCompareRequest() *CompareRequest {
	o := &CompareRequest{
		Priority: defualtComparePriority,
		Sources:  make([]*CompareSource, 0),
	}
	return o
}
