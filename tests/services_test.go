package tests

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	"github.com/chnykn/bimface/bean/request"
	"github.com/chnykn/bimface/utils"
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

func fileToBuffer(file io.Reader) (*bytes.Buffer, error) {

	result := bytes.NewBuffer(make([]byte, 0))

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if (err != nil) && (err != io.EOF) {
			return nil, err
		}

		if n == 0 {
			break
		}

		_, err = result.Write(buf[:n])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func TestUploadService(t *testing.T) {
	client := getClient()

	filePath := "F:/Garry/RVT/Sample.rvt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("err = %v \n", err)
		return
	}
	defer file.Close()

	buffer, err := fileToBuffer(file)
	if err != nil {
		fmt.Printf("err = %v \n", err)
		return
	}

	fileName := path.Base(filePath)

	uploadReq := &request.UploadRequest{
		Name:   utils.EncodeURI(fileName),
		Buffer: buffer,
	}

	fileBean, err := client.UploadService.Upload(uploadReq)
	if err == nil {
		fmt.Printf("Upload = %v \n", fileBean)
	} else {
		fmt.Printf("err = %v \n", err)
	}

	/*
		fileBean, err := client.UploadService.GetFileMetadata(1538038887260832)
		if err == nil {
			fmt.Printf("GetFileMetadata = %v \n", fileBean)
		} else {
			fmt.Printf("err = %v \n", err)
		}

		result, err := client.UploadService.DeleteFile(1538038887260832)
		if err == nil {
			fmt.Printf("DeleteFile = %v \n", result)
		} else {
			fmt.Printf("err = %v \n", err)
		}
	*/

}

func TestIntegrateService(t *testing.T) {
	client := getClient()

	status, err := client.IntegrateService.GetIntegrateStatus(1538044326628288)
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

func TestSupportFileService(t *testing.T) {
	client := getClient()

	res, err := client.SupportFileService.GetSupport()
	if err == nil {
		fmt.Printf("GetSupport = %v \n", res)
	} else {
		fmt.Printf("err = %v \n", err)
	}

}

//

func TestUrlEscape(t *testing.T) {

	s := "doc/callback?name=trans&id=1&fileId=1602047233139584&transferId=1602047233139584&status=success&thumbnail=https%3A%2F%2Fm.bimface.com%2F0c18dfb2c6544f15e88e9a6d5717f8aa%2Fthumbnail%2F96.png%2Chttps%3A%2F%2Fm.bimface.com%2F0c18dfb2c6544f15e88e9a6d5717f8aa%2Fthumbnail%2F256.png&reason=&nonce=785eb7a6-5446-4f40-bbf7-5e0d6c0b307a&signature=6b7564994a6f2b24a51f5165cac3284c"

	res, err := utils.DecodeURI(s)
	if err != nil {
		fmt.Printf("DecodeURI err=%v \n", err)
	}
	fmt.Println(res)

}
