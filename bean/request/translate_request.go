package request

//TranslateSource ***
type TranslateSource struct {
	FileID     int64  //必填
	Compressed bool   //是否为压缩文件 默认为false
	RootName   string //如果是压缩文件，必须指定压缩包中哪一个是主文件
}

//NewTranslateSource ***
func NewTranslateSource(fileID int64) *TranslateSource {
	o := &TranslateSource{
		FileID:     fileID,
		Compressed: false,
		RootName:   "",
	}
	return o
}

//---------------------------------------------------------------------

const (
	defualtTranslatePriority int = 2
)

//TranslateRequest ***
type TranslateRequest struct {
	Source   *TranslateSource
	Priority int //[1,2,3] 数字越大，优先级越低
	Callback string
}

//NewTranslateRequest ***
func NewTranslateRequest(sourceFileID int64) *TranslateRequest {
	o := &TranslateRequest{
		Source:   NewTranslateSource(sourceFileID),
		Priority: defualtTranslatePriority,
		Callback: "",
	}
	return o
}
