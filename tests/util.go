package tests

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveToJsonFile(obj any, jsonfile string) {
	file, err := os.Create(jsonfile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t") // 设置缩进和前缀
	err = encoder.Encode(obj)

	if err != nil {
		fmt.Println("Error encoding person:", err)
	}
}
