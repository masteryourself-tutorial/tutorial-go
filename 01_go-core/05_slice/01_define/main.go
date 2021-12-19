package main

import "fmt"

func main() {
	sliceDefine1()
	sliceDefine2()
	sliceDefine3()
}

// 让切片去引用一个已经创建好的数组
func sliceDefine1() {
	var array = [5]int{1, 2, 3, 4, 5}
	var slice = array[1:3]
	// [2 3]
	fmt.Println(slice)
}

// 通过 make 创建切片
func sliceDefine2() {
	// 第一个参数是数据类型，第二个参数是长度，第三个参数是容量
	var slice = make([]int32, 5, 10)
	// slice 的长度是 5, 容量是 10
	fmt.Printf("slice 的长度是 %v, 容量是 %v \n", len(slice), cap(slice))
}

// 定义切片并直接赋值
func sliceDefine3() {
	var slice = []int32{10, 20, 30}
	// slice 的长度是 3, 容量是 3
	fmt.Printf("slice 的长度是 %v, 容量是 %v \n", len(slice), cap(slice))
}
