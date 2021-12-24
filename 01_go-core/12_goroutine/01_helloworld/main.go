package main

import (
	"fmt"
	"time"
)

// 使用 go 开启一个协程
func main() {
	go helloWorld()
	for i := 0; i < 10; i++ {
		fmt.Println("hello main", i)
		time.Sleep(time.Second)
	}
}

func helloWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world", i)
		time.Sleep(time.Second)
	}
}
