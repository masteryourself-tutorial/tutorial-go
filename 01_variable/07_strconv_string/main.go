package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 = "true"
	var n2 = "100"
	var n3 = "43.99"

	v1, _ := strconv.ParseBool(n1)
	fmt.Printf("v1 值是 %v, 类型是 %T \n", v1, v1)

	// 第二个参数表示进制, 第三个表示转成 int32
	v2, _ := strconv.ParseInt(n2, 10, 32)
	fmt.Printf("v2 值是 %v, 类型是 %T \n", v2, v2)

	// 第二个参数表示转成 float64
	v3, _ := strconv.ParseFloat(n3, 64)
	fmt.Printf("v3 值是 %v, 类型是 %T \n", v3, v3)
}
