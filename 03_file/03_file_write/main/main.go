package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "03_file/03_file_write/example.txt"
	// 只写和创建
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	// 只写和追加
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	// 延时关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件关闭失败", err)
		}
	}(file)
	// 使用带缓存的 buffer writer
	writer := bufio.NewWriter(file)
	var str = "hello"
	res, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("文件写入失败", err)
	}
	fmt.Printf("写入 %v 个字符 \n", res)
	// 因为是带缓存的, 所以是暂存在内存中, 所以调用 flush() 写入到磁盘中
	err = writer.Flush()
	if err != nil {
		fmt.Println("文件刷新失败", err)
	}
}
