package response

type IntegrateFileData struct {
	IntegrateId int64 `json:"integrateId"`

	FileId   int64  `json:"fileId"`
	FileName string `json:"fileName"`
	LinkedBy string `json:"linkedBy"`

	DatabagId         string `json:"databagId"`
	DrawingSheetCount int    `json:"drawingSheetCount"`

	Floor     string  `json:"floor"`
	FloorSort float64 `json:"floorSort"`

	Specialty     string  `json:"specialty"`
	SpecialtySort float64 `json:"specialtySort"`
}

type IntegrateFileBean IntegrateFileData
