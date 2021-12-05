package main

import (
	"fmt"
)

func main() {
	var num int32 = 998

	// 定义一个指针变量
	var ptr *int32 = &num
	// 打印
	fmt.Printf("ptr 是 %v \n", ptr)
	fmt.Printf("ptr 的地址是 %v \n", &ptr)
	fmt.Printf("ptr 地址对应的值是 %v \n", *ptr)
}
