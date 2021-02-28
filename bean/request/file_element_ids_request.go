package request

type FileIdHashWithElementIds struct {
	FileIdHash string   `json:"fileIdHash"`
	ElementIds []string `json:"elementIds"`
}

type FileElementIdsRequest []*FileIdHashWithElementIds
