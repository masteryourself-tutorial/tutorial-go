package main

import (
	"fmt"
)

func main()  {
	test()
	fmt.Println("test 之后的代码")
}

func test()  {
	// 使用 defer + recover 捕获异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获异常", err)
		}
	}()
	num1 := 10
	num2 := 5
	res := num1 / (5 - num2)
	fmt.Println("res := ", res)
}