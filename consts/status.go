// Copyright 2019-2021 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package consts

//Status Status
type Status string

const (
	//SUCCESS for CallbackStatus and TransferStatus
	SUCCESS Status = "success"

	//FAILED for CallbackStatus and TransferStatus
	FAILED Status = "failed"

	//PREPARE for TransferStatus
	PREPARE Status = "prepare"

	//PROCESSING for TransferStatus
	PROCESSING Status = "processing"
)
