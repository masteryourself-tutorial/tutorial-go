package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	fmt.Println("test 之后的代码")

	err := test2("error")
	if err != nil {
		// 直接终止程序
		panic(err)
	} else {
		fmt.Println("正常执行了")
	}
	fmt.Println("test2 之后的代码")
}

func test() {
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

func test2(name string) (err error) {
	if name == "error" {
		return errors.New("入参不正确")
	} else {
		return nil
	}
}
