package main

import (
	"errors"
	"fmt"
)

func main() {
	err := testError(-8)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Printf("程序继续执行")
}

func test() {
	// 使用 defer + recover 来处理异常
	defer func() {
		// 使用 recover 可以捕获异常
		err := recover()
		if err != nil {
			fmt.Println("程序异常了", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res = ", res)
}

// 使用 errors.New 返回自定义异常
func testError(i int) (err error) {
	if i < 0 {
		return errors.New("参数不能小于0")
	}
	return nil
}
