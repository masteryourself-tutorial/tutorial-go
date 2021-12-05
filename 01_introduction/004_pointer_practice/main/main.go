package main

import (
	"fmt"
)

func main() {

	modifyPointer()

	fmt.Println("=============================== 分隔符 ===============================")

	modifyValue()
}

func modifyPointer() {
	// 修改前
	var num int32 = 9527
	var ptr *int32 = &num
	fmt.Printf("修改前 num 的内存地址是 %v, 值是 %v \n", &num, num)
	fmt.Printf("修改前 ptr 的内存地址是 %v, 值是 %v \n", &ptr, *ptr)

	// 修改指针地址(*ptr 表示 ptr 的内存地址)对应的值
	// ptr 内存地址不变, 内存地址指向的值发生变化
	// num 内存地址不变, 内存地址指向的值发生变化
	*ptr = 10

	// 修改后
	fmt.Printf("修改后 num 的内存地址是 %v, 值是 %v \n", &num, num)
	fmt.Printf("修改后 ptr 的内存地址是 %v, 值是 %v \n", &ptr, *ptr)
}

func modifyValue() {
	// 修改前
	var num int32 = 9527
	var ptr *int32 = &num

	fmt.Printf("修改前 num 的内存地址是 %v, 值是 %v \n", &num, num)
	fmt.Printf("修改前 ptr 的内存地址是 %v, 值是 %v \n", &ptr, *ptr)

	// 修改指针
	// ptr 的内存地址不变, 但值发生了变化, 因为重新指向了
	// num 的内存地址和值均不变
	var num2 int32 = 10
	ptr = &num2

	// 修改后
	fmt.Printf("修改后 num 的内存地址是 %v, 值是 %v \n", &num, num)
	fmt.Printf("修改后 ptr 的内存地址是 %v, 值是 %v \n", &ptr, *ptr)
}
