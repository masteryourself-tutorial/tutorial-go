package main

import "fmt"

func main() {
	var v1 = 1
	var v2 = 123.456
	var v3 = false
	var v4 = 'c'
	var str string

	str = fmt.Sprintf("%d", v1)
	fmt.Printf("str 值是 %v, 类型是 %T \n", str, str)

	str = fmt.Sprintf("%f", v2)
	fmt.Printf("str 值是 %v, 类型是 %T \n", str, str)

	str = fmt.Sprintf("%t", v3)
	fmt.Printf("str 值是 %v, 类型是 %T \n", str, str)

	str = fmt.Sprintf("%c", v4)
	fmt.Printf("str 值是 %v, 类型是 %T \n", str, str)
}
