package main

import "fmt"

func main() {
	addFunc := add()
	// 11
	fmt.Print(addFunc(1))
}

// 返回值是一个函数
func add() func(int) int {
	var n = 10
	// 形成闭包
	return func(x int) int {
		return n + x
	}
}
