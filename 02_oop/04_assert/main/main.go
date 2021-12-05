package main

import "fmt"

func main() {
	var a int32 = 18
	var b interface{}
	b = a
	if a, ok := b.(int64); ok {
		fmt.Printf("转换成功, a 的类型是 %T, a 的值是 %v \n", a, a)
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println("程序结束")
}
