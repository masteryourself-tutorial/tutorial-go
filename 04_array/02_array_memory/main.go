package main

import "fmt"

// int 默认是 int64, 占用 8 个字节
func main() {
	var array = [3]int{100, 200, 300}
	// array 的首地址是 0xc0000b2000
	fmt.Printf("array 的首地址是 %p \n", &array)
	// array[1] 的首地址是 0xc0000b2000, array[2] 的首地址是 0xc0000b2008, array[3] 的首地址是 0xc0000b2010
	fmt.Printf("array[1] 的首地址是 %p, array[2] 的首地址是 %p, array[3] 的首地址是 %p \n",
		&array[0], &array[1], &array[2])
}
