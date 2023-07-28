package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestGetDB(t *testing.T) {

	filePath := "../public/76c07abe-fd3f-445d-aaf5-fb1502569e75.mp4"

	// 检查文件是否存在
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("File does not exist:", filePath)
			fmt.Println("File does not exist:", filePath)
		} else {
			log.Info("Error checking file:", err)
			fmt.Println("Error checking file:", err)
		}
		return
	}

	// 检查文件是否具有读权限
	file, err := os.Open(filePath)
	if err != nil {
		log.Info("Error opening file:", err)
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 文件具有读权限
	fmt.Println("File exists and has read permission:", filePath)
}
