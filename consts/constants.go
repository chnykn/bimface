package consts

import "time"

const (
	//Charset UTF_8 = Charset.forName("UTF-8");
	//MediaType MEDIA_TYPE_JSON = MediaType.parse("application/json; charset=utf-8");

	//APIHost ***
	APIHost string = "https://api.bimface.com"

	//FileHost ***
	FileHost string = "https://file.bimface.com"

	//StreamMime ***
	StreamMime string = "application/octet-stream"

	//JSONMime ***
	JSONMime string = "application/json"

	//FormMime ***
	FormMime string = "application/x-www-form-urlencoded"

	//BlockSize ***
	BlockSize int = 4194304

	//PutThreshold ***
	PutThreshold int = 4194304

	//DefaultMaxIdleConnections ***
	DefaultMaxIdleConnections int = 32

	//DefaultKeepAliveDurationNs ***
	DefaultKeepAliveDurationNs int = 300000

	//DefaultMaxRequests ***
	DefaultMaxRequests int = 64

	//DefaultMaxRequestsPerHost ***
	DefaultMaxRequestsPerHost int = 5

	//DefaultConnectTimeout ***
	DefaultConnectTimeout time.Duration = 20 * time.Second

	//DefaultWriteTimeout ***
	DefaultWriteTimeout int = 0

	//DefaultResponseTimeout ***
	DefaultResponseTimeout int = 30
)
