// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package request

type FileIdWithElementIds struct {
	FileId     string   `json:"fileIdHash"`
	ElementIds []string `json:"elementIds"`
}

type FileElementIdsRequest []*FileIdWithElementIds
