package main

import "fmt"

// 把函数赋值给变量, 通过变量调用方法
func main() {
	addVariable := add
	// addVariable 的类型是 func(int, int) int, add() 的类型是 func(int, int) int
	fmt.Printf("addVariable 的类型是 %T, add() 的类型是 %T \n", addVariable, add)
	res := addVariable(10, 20)
	// 通过变量调用函数结果是 30
	fmt.Printf("通过变量调用函数结果是 %v \n", res)
}

func add(x1 int, x2 int) int {
	return x1 + x2
}
