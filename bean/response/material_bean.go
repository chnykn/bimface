package response

import "github.com/chnykn/bimface/v2/bean/common"

type MaterialInfo struct {
	Id         string                `json:"id"`
	Name       string                `json:"name"`
	Parameters *common.PropertyGroup `json:"parameters,omitempty"`
}

type MaterialBean MaterialInfo
