package service

import (
	"bimface/config"
	"bimface/http"
)

//AbstractService ***
type AbstractService struct {
	Endpoint      *config.Endpoint
	ServiceClient *http.ServiceClient
}
