package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	openFileWithBuf()
	openFileOnce()
}

// 使用带 buff 的缓冲区读取
func openFileWithBuf() {
	// 打开文件
	f, err := os.Open("/Users/ruanrenzhao/go_project/tutorial-go/01_go-core/08_file/doc/file.txt")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	// 延时退出
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("文件关闭失败", err)
		}
	}(f)
	// 使用 bufio 读取内容, 带缓冲区的, 默认是 4096 字节
	reader := bufio.NewReader(f)
	buf := make([]byte, 4096)
	// 打印内容
	for {
		n, err := reader.Read(buf)
		if err != nil {
			//遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return
			}
			fmt.Println("文件读取失败", err)
		}
		fmt.Println("读取内容:", string(buf[:n]))
	}
}

// 一次性打开文件, 适用于文件内容比较小的场景
func openFileOnce() {
	// 不需要打开文件, 也不需要关闭文件, 函数已经封装好了
	content, err := ioutil.ReadFile("/Users/ruanrenzhao/go_project/tutorial-go/01_go-core/08_file/doc/file.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	fmt.Printf("文件内容是 %v \n", string(content))
}
