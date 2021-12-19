package main

import "fmt"

func main() {
	arrayDefine1()
	arrayDefine2()
	arrayDefine3()
}

// 数组定义方式 1
func arrayDefine1() {
	var array [3]int = [3]int{100, 200, 300}
	fmt.Println(array)
}

// 数组定义方式 2: [...] 规定写法
func arrayDefine2() {
	var array = [...]int{100, 200, 300}
	for i := 0; i < len(array); i++ {
		fmt.Printf("index = %v, value = %v \n", i, array[i])
	}
}

// 数组定义方式 3, 按照下标赋值
func arrayDefine3() {
	var array = [3]int{2: 300, 1: 200, 0: 100}
	for index, value := range array {
		fmt.Printf("index = %v, value = %v \n", index, value)
	}
}
