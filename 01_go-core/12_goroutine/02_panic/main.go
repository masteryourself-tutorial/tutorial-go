package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	time.Sleep(time.Second)
	fmt.Println("程勋运行结束")
}

// 使用 recover 捕获异常
func test() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("程序异常了", err)
		}
	}()
	a := 100
	b := 0
	fmt.Println("协程执行到这里")
	return a / b
}
