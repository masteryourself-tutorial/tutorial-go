package main

import "fmt"

func main() {
	res := add(10, 20)
	// 32
	fmt.Println("res = ", res)
}

func add(x1 int, x2 int) int {
	// 10
	defer fmt.Println("defer x1 值是", x1)
	// 20
	defer fmt.Println("defer x2 值是", x2)
	x1++
	x2++
	// 11
	fmt.Println("x1 值是", x1)
	// 21
	fmt.Println("x2 值是", x2)
	return x1 + x2
}
