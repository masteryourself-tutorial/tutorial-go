package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/Users/ruanrenzhao/go_project/tutorial-go/01_go-core/08_file/doc/file.txt"
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("文件或者目录不存在")
		return
	}
	if err == nil {
		fmt.Println("文件或者目录存在")
	}
	fmt.Println("是否是文件夹", stat.IsDir())
}
