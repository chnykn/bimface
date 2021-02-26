package response

/*
   "failedReason" : "input.stream.read.error",
   "fileId" : 1216113551663296,
   "name" : "-1F.rvt",
   "status" : "failure"
*/

//UploadStatus ***
type UploadStatus struct {
	FileId       int64  `json:"fileId"`
	Name         string `json:"name"`
	FailedReason int64  `json:"failedReason"`
	Status       string `json:"status"`
}
