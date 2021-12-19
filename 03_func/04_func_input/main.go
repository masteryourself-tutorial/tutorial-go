package main

import "fmt"

// 函数作为形参调用
func main() {
	res := funcInput(add, 10, 20)
	fmt.Printf("函数形参调用结果是 %v \n", res)
}

func funcInput(addFunc func(int, int) int, x1 int, x2 int) int {
	return addFunc(x1, x2)
}

func add(x1 int, x2 int) int {
	return x1 + x2
}
