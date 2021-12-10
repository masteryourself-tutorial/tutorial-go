package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 使用带缓冲区的 buff 读取内容
	fmt.Printf("\n==================== 分隔符: openFileWithBuf ====================\n")
	openFileWithBuf()
	// 使用 ioutil 一次性读取内容
	fmt.Printf("\n==================== 分隔符: openFileOnce ====================\n")
	openFileOnce()
}

func openFileWithBuf() {
	// 打开文件
	f, err := os.Open("/Users/ruanrenzhao/go_project/tutorial-go/README.md")
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
	// 打印内容
	for {
		if content, err := reader.ReadString('\n'); err != io.EOF {
			fmt.Print("读取内容 》》》", content)
		} else {
			fmt.Println("文件读完")
			break
		}
	}
}

func openFileOnce() {
	// 一次性打开文件, 适用于文件内容比较小的场景
	// 不需要打开文件, 也不需要关闭文件
	content, err := ioutil.ReadFile("/Users/ruanrenzhao/go_project/tutorial-go/README.md")
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	fmt.Printf("文件内容是 %v \n", string(content))
}
