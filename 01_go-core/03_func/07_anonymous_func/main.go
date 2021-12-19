package main

import "fmt"

// 3. 全局匿名函数，如果把一个函数赋值给一个全局变量，那么这个匿名函数就成为一个全局匿名函数，整个程序有效
var (
	AddFunc = func(x int, y int) int {
		return x + y
	}
)

func main() {
	anonymousFunc1()
	anonymousFunc2()
	fmt.Println("res = ", AddFunc(1, 1))
	fmt.Println("res = ", AddFunc(2, 3))
}

// 1. 定义时就直接使用，这种只能调用一次
func anonymousFunc1() {
	res := func(x int, y int) int {
		return x + y
	}(10, 20)
	fmt.Println("res = ", res)
}

// 2. 将匿名函数赋值给一个变量，通过变量调用，这种可以调用多次
func anonymousFunc2() {
	addFunc := func(x int, y int) int {
		return x + y
	}
	fmt.Println("res = ", addFunc(10, 20))
	fmt.Println("res = ", addFunc(20, 20))
}
