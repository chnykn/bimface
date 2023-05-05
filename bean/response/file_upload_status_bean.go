// Copyright 2019-2023 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type FileUploadStatusBean struct {
	FileId       int64  `json:"fileId"`
	Name         string `json:"name"`
	FailedReason int64  `json:"failedReason"`
	Status       string `json:"status"`
}
