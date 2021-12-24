package main

import (
	"fmt"
	"time"
)

// 定义一个 dataChan 用来存储数据
var dataChan = make(chan int, 20)

// 定义一个 flagChan 用来表示数据读取结束
var flagChan = make(chan bool, 1)

func main() {
	// 开启一个协程写数据
	go writeData()
	// 开启一个协程读数据
	go readData()
	// 从 flagChan 中读取到数据才认为结束
	if _, ok := <-flagChan; !ok {
	}
	fmt.Println("任务结束")
}

func writeData() {
	for i := 1; i <= cap(dataChan); i++ {
		dataChan <- i
		fmt.Printf("<<<<<< 放入数据 %v \n", i)
		time.Sleep(time.Millisecond * 50)
	}
	// 调用 close 方法
	close(dataChan)
}

func readData() {
	for {
		value, ok := <-dataChan
		// 说明管道已经关闭了
		if !ok {
			break
		}
		fmt.Printf(">>> dataChan 获取数据 %v \n", value)
		time.Sleep(time.Millisecond * 100)
	}
	// 如果全部读取完毕, 向 flagChan 中写入数据
	flagChan <- true
	close(flagChan)
}
