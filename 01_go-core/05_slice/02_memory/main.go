package main

import "fmt"

func main() {
	array := [5]int{100, 200, 300, 400, 500}
	// array[0:2] 表示起始下标为 0, 结束下标为 2（但是不包含 2）
	slice := array[0:2]
	// slice 的长度是 2, 容量是 5
	fmt.Printf("slice 的长度是 %v, 容量是 %v \n", len(slice), cap(slice))
	// array 的首地址是 0xc000092030
	fmt.Printf("array 的首地址是 %p \n", &array)
	// slice 的地址是 0xc000114000
	fmt.Printf("slice 的地址是 %p \n", &slice)
	// slice[0] 的首地址是 0xc000092030, slice[1] 的首地址是 0xc000092038
	fmt.Printf("slice[0] 的首地址是 %p, slice[1] 的首地址是 %p \n", &slice[0], &slice[1])
}
