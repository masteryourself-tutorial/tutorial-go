package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("\n==================== 分隔符: copyReadAndWrite ====================\n")
	copyReadAndWrite()
	fmt.Printf("\n==================== 分隔符: ioCopy ====================\n")
	ioCopy()
}

func copyReadAndWrite() {
	sourceFilePath := "03_file/04_file_copy/doc/source/hello.txt"
	destFilePath := "03_file/04_file_copy/doc/dest/hello.txt"
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

func ioCopy() {
	sourceFile, err := os.Open("03_file/04_file_copy/doc/source/风景.jpeg")
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
	destFile, err := os.OpenFile("03_file/04_file_copy/doc/dest/风景.jpeg", os.O_WRONLY|os.O_CREATE, 0666)
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
