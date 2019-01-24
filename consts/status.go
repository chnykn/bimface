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
