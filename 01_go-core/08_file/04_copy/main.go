package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	copyReadAndWrite()
	ioCopy()
}

// 读取到 buffer 中再将 buffer 写入到文件里
func copyReadAndWrite() {
	sourceFilePath := "/Users/ruanrenzhao/go_project/tutorial-go/01_go-core/08_file/doc/file.txt"
	destFilePath := "/Users/ruanrenzhao/go_project/tutorial-go/01_go-core/08_file/doc/file_copy.txt"
	data, err := ioutil.ReadFile(sourceFilePath)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	err = ioutil.WriteFile(destFilePath, data, 0666)
	if err != nil {
		fmt.Println("文件写入失败", err)
	}
}

// 使用 io.Copy() 工具类
func ioCopy() {
	sourceFile, err := os.Open("/Users/ruanrenzhao/go_project/tutorial-go/01_go-core/08_file/doc/树叶.jpeg")
	if err != nil {
		fmt.Println("源文件打开失败", err)
		return
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			fmt.Println("文件关闭失败", err)
		}
	}(sourceFile)
	destFile, err := os.OpenFile("/Users/ruanrenzhao/go_project/tutorial-go/01_go-core/08_file/doc/树叶_copy.jpeg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("目标文件打开失败", err)
		return
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			fmt.Println("文件关闭失败", err)
		}
	}(destFile)
	res, err := io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println("文件 copy 失败")
	}
	fmt.Printf("写入字节总数 %v \n", res)
}
