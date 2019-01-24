package service

import (
	"bimface/bean/response"
	"bimface/config"
	"bimface/http"
	"bimface/utils"
)

const supportFileURI = "/support"

var defSupportFile = &response.SupportFile{
	Length: 13,
	Types:  []string{"rvt", "rfa", "dwg", "dxf", "skp", "ifc", "dgn", "obj", "stl", "3ds", "dae", "ply", "igms"},
}

//SupportFileService ***
type SupportFileService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewSupportFileService ***
func NewSupportFileService(serviceClient *http.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *SupportFileService {
	o := &SupportFileService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //http.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *SupportFileService) supportFileURL() string {
	return o.Endpoint.FileHost + supportFileURI
}

//GetSupportWithAccessToken ***
func (o *SupportFileService) GetSupportWithAccessToken(token string) (*response.SupportFile, *utils.Error) {
	/*
		headers := http.NewHeaders()
		headers.AddOAuth2Header(token)

		resp := o.ServiceClient.Get(o.supportFileURL(), headers.Header)

		result := response.NewSupportFile()
		err := http.RespToBean(resp, result)

		return result, err
	*/
	return defSupportFile, nil
}

//GetSupport ***
func (o *SupportFileService) GetSupport() (*response.SupportFile, *utils.Error) {
	/*
		accessToken, err := o.AccessTokenService.Get()
		if err != nil {
			return nil, err
		}

		return o.GetSupportWithAccessToken(accessToken.Token)
	*/
	return defSupportFile, nil
}
