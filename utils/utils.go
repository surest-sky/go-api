package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func StoragePath(dir string) string {
	return "storage/" + dir
}

func UploadPath() string {
	path := StoragePath("resource")
	path = filepath.Join(path, time.Now().Format("20060102"))
	CreateDateDir(path)
	return path
}

func CreateDateDir(basePath string){
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		err = os.MkdirAll(basePath, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
		// 再修改权限
		err = os.Chmod(basePath, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
