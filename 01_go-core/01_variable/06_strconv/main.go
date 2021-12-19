package main

import (
	"fmt"
	"strconv"
)

func main() {
	var v1 = 1
	var v2 = 123.456
	var v3 = false
	var str string

	// 第一个参数只接收 int64 位, 第二个参数表示进制
	str = strconv.FormatInt(int64(v1), 10)
	fmt.Printf("str 值是 %v, 类型是 %T \n", str, str)

	// 第二个参数表示格式, 第三个参数表示尾数部分精度, 第四个参数表示来源是 float64
	str = strconv.FormatFloat(v2, 'f', 10, 64)
	fmt.Printf("str 值是 %v, 类型是 %T \n", str, str)

	str = strconv.FormatBool(v3)
	fmt.Printf("str 值是 %v, 类型是 %T \n", str, str)
}
