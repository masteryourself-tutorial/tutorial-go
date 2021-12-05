package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	f, err := os.Open("/Users/ruanrenzhao/go_project/tutorial-go/README.md")
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 延时退出
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("文件读取失败", err)
		}
	}(f)
	// 使用 bufio 读取内容
	reader := bufio.NewReader(f)
	// 打印内容
	for {
		if content, err := reader.ReadString('\n'); err != io.EOF {
			fmt.Println("读取内容 》》》", content)
		} else {
			fmt.Println("文件读完")
			break
		}
	}
}
