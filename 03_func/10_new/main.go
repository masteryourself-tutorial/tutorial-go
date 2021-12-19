package main

import "fmt"

// 使用 new 分配内存
func main() {
	n := new(int)
	*n = 10
	// n 的类型是 *int, n 的地址是 0xc00000e028, n 的值是 0xc00001c070，n 地址指向的值是 10
	fmt.Printf("n 的类型是 %T, n 的地址是 %v, n 的值是 %v，"+
		"n 地址指向的值是 %v \n", n, &n, n, *n)
}
