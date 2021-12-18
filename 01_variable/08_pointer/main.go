package main

import "fmt"

func main() {
	var i int = 30
	// i 的地址是 0xc0000b0008, 值是 30
	fmt.Printf("i 的地址是 %v, 值是 %v \n", &i, i)
	// j 是一个指针变量, 类型是 *int, 本身的值是个地址 &i
	var j *int = &i
	// j 的地址是 0xc0000aa020, 值是 0xc0000b0008, 地址指向的值是 30
	fmt.Printf("j 的地址是 %v, 值是 %v, 地址指向的值是 %v \n", &j, j, *j)
}
