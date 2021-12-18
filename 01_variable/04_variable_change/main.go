package main

import "fmt"

func main() {
	var x int8 = 32
	y := float32(x)
	z := int64(x)
	// x = 32 类型是 int8, y = float32 类型是 float32, z = int64 类型是 int64
	fmt.Printf("x = %v 类型是 %T, y = %T 类型是 %T, z = %T 类型是 %T \n",
		x, x, y, y, z, z)
}
