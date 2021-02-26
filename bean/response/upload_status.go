package response

type FileUploadStatusBean struct {
	FileId       int64  `json:"fileId"`
	Name         string `json:"name"`
	FailedReason int64  `json:"failedReason"`
	Status       string `json:"status"`
}
