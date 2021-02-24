package request

//TranslateSource ***
type TranslateSource struct {
	FileId     int64  `json:"fileId"`               //必填
	Compressed bool   `json:"compressed,omitempty"` //是否为压缩文件 默认为false
	RootName   string `json:"rootName,omitempty"`   //如果是压缩文件，必须指定压缩包中哪一个是主文件
}

//NewTranslateSource ***
func NewTranslateSource(fileId int64) *TranslateSource {
	o := &TranslateSource{
		FileId:     fileId,
		Compressed: false,
		RootName:   "",
	}
	return o
}

//-----------------------

//TranslateRequest ***
type TranslateRequest struct {
	Source *TranslateSource `json:"source"`

	//转换引擎自定义参数，config参数跟转换引擎相关，不同的转换引擎支持不同的config格式。
	//例如转换时添加内置材质，则添加参数值{"texture":true}，添加外部材质时参考“使用模型外置材质场景”请求报文
	//详见：http://static.bimface.com/restful-apidoc/dist/translateSingleModel.html#_translateusingput_1
	Config interface{} `json:"config,omitempty"` //Json Object

	Callback string `json:"callback,omitempty"`
}

//NewTranslateRequest ***
func NewTranslateRequest(sourceFileId int64) *TranslateRequest {
	o := &TranslateRequest{
		Source:   NewTranslateSource(sourceFileId),
		Config:   nil,
		Callback: "",
	}
	return o
}
