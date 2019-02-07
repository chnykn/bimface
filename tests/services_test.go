package tests

import (
	"fmt"
	"testing"

	"github.com/chnykn/bimface/bean/request"
)

func TestAccessTokenService(t *testing.T) {
	client := getClient()
	accessToken, err := client.AccessTokenService.Get()
	if err == nil {
		fmt.Printf("AccessTokenService.Get() = %v \n", accessToken)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}

func TestDownloadService(t *testing.T) {
	client := getClient()
	url, err := client.DownloadService.GetDownloadURL(1531158346498720, "")
	if err == nil {
		fmt.Printf("GetDownloadURL = %v \n", url)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}

func TestCategoryTreeService(t *testing.T) {
	client := getClient()

	cates, err := client.CategoryTreeService.GetCategoryTree(1525914374062720)
	if err == nil {
		fmt.Printf("GetCategoryTree = %v \n", cates)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	tree, err := client.CategoryTreeService.GetCategoryTreeV2(1525914374062720)
	if err == nil {
		fmt.Printf("GetCategoryTreeV2 = %v \n", tree)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}

func TestIntgrTreeService(t *testing.T) {
	client := getClient()

	stree, err := client.IntgrTreeService.GetSpecialtyTree(1532041801780160)
	if err == nil {
		fmt.Printf("GetSpecialtyTree = %v \n", stree)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	ftree, err := client.IntgrTreeService.GetFloorTree(1532041801780160)
	if err == nil {
		fmt.Printf("GetFloorTree = %v \n", ftree)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}

func TestElementService(t *testing.T) {
	client := getClient()

	elemIds, err := client.ElementService.GetElementsWithParams(1525914374062720, nil)
	if err == nil {
		fmt.Printf("GetElements = %v \n", elemIds)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	elems, err := client.ElementService.GetIntgrElementsWithParams(1532041801780160, nil)
	if err == nil {
		fmt.Printf("GetIntgrElementsWithParams = %v \n", elems)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}

func TestFloorService(t *testing.T) {
	client := getClient()

	Floors, err := client.FloorService.GetFloors(1525914374062720)
	if err == nil {
		fmt.Printf("GetFloors = %v \n", Floors)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	Floors, err = client.FloorService.GetIntegrationFloors(1532041801780160)
	if err == nil {
		fmt.Printf("GetIntegrationFloors = %v \n", Floors)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}

func TestPropertyService(t *testing.T) {
	client := getClient()

	props, err := client.PropertyService.GetElementProperty(1525914374062720, "678602")
	if err == nil {
		fmt.Printf("GetElementProperty = %v \n", props)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	props, err = client.PropertyService.GetIntgrElementProperty(1532041801780160, 1525914374062720, "678602")
	if err == nil {
		fmt.Printf("GetIntgrElementProperty = %v \n", props)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}

func TestTranslateService(t *testing.T) {
	client := getClient()

	transRequest := request.NewTranslateRequest(1525914374062720)
	status, err := client.TranslateService.Translate(transRequest)
	if err == nil {
		fmt.Printf("Translate = %v \n", status)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	status, err = client.TranslateService.GetTranslateStatus(1525914374062720)
	if err == nil {
		fmt.Printf("GetTranslateStatus = %v \n", status)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}

func TestShareLinkService(t *testing.T) {
	client := getClient()

	shareLink, err := client.ShareLinkService.CreateShare(1531158346498720, 0)
	if err == nil {
		fmt.Printf("CreateShare = %v \n", shareLink)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	result, err := client.ShareLinkService.DeleteShare(1531158346498720)
	if err == nil {
		fmt.Printf("DeleteShare = %v \n", result)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}

func TestUploadService(t *testing.T) {
	client := getClient()

	fileBean, err := client.UploadService.GetFileMetadata(1538038887260832)
	if err == nil {
		fmt.Printf("GetFileMetadata = %v \n", fileBean)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	/*
		result, err := client.UploadService.DeleteFile(1538038887260832)
		if err == nil {
			fmt.Printf("DeleteFile = %v \n", result)
		} else {
			fmt.Printf("err = %v \n", err)
		}

		uploadReq := request.NewUploadRequest("F:/场地_建筑.rvt")
		fileBean, err := client.UploadService.Upload(uploadReq)
		if err == nil {
			fmt.Printf("Upload = %v \n", fileBean)
		} else {
			fmt.Printf("err = %v \n", err)
		}
	*/

}

func TestIntegrateService(t *testing.T) {
	client := getClient()

	status, err := client.IntegrateService.GetIntgrStatus(1538044326628288)
	if err == nil {
		fmt.Printf("GetIntgrStatus = %v \n", status)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	integReq := request.NewIntgrRequest("深高速-综合办公楼-New")
	integReq.Sources = append(integReq.Sources, request.NewIntgrSource(1531158346498720, "建筑"))
	integReq.Sources = append(integReq.Sources, request.NewIntgrSource(1525914374062720, "结构"))
	status, err = client.IntegrateService.Integrate(integReq)
	if err == nil {
		fmt.Printf("Integrate = %v \n", status)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	result, err := client.IntegrateService.DeleteIntegrate(1538044326628288)
	if err == nil {
		fmt.Printf("DeleteIntegrate = %v \n", result)
	} else {
		fmt.Printf("err = %v \n", err)
	}
}
