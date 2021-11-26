package main

import (
	"fmt"
)

func main()  {
	num1 := 10
	fmt.Printf("num1 的类型是 %T, num1 的值是 %v, num1 的地址是 %v\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2 的类型是 %T, num2 的值是 %v, num2 的地址是 %v, num2 指针地址对应的值是 %v\n", 
	num2, num2, &num2, *num2)

	*num2 = 10
	fmt.Printf("num2 的类型是 %T, num2 的值是 %v, num2 的地址是 %v, num2 指针地址对应的值是 %v\n", 
	num2, num2, &num2, *num2)
}