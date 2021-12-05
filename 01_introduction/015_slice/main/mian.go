package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n==================== 分隔符: demo1 ====================\n")
	demo1()
	fmt.Printf("\n==================== 分隔符: demo2 ====================\n")
	demo2()
	fmt.Printf("\n==================== 分隔符: demo3 ====================\n")
	demo3()
	fmt.Printf("\n==================== 分隔符: demo4 ====================\n")
	demo4()
	fmt.Printf("\n==================== 分隔符: demo5 ====================\n")
	demo5()
	fmt.Printf("\n==================== 分隔符: demo6 ====================\n")
	demo6()
	fmt.Printf("\n==================== 分隔符: demo7 ====================\n")
	demo7()
	fmt.Printf("\n==================== 分隔符: demo8 ====================\n")
	demo8()
}

// 定义方式一: 使用数组定义切片
func demo1() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// 从数组下标[1,3), 左闭右开, 所以为 [2, 3]
	var slice = arr[1:3]
	fmt.Printf("slice 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice, len(slice), cap(slice))
}

// 切片是引用类型
func demo2() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var slice = arr[1:3]
	fmt.Printf("修改前: arr 数组是 %v \n", arr)
	fmt.Printf("修改前: slice 切片是 %v \n", slice)
	slice[1] = 111
	// slice 是引用类型, 从底层结构来说是一个 struct 结构体, 它改变会影响数组的值
	fmt.Printf("修改后: arr 数组是 %v \n", arr)
	fmt.Printf("修改后: slice 切片是 %v \n", slice)
}

// 定义方式二: 使用 make 定义切片
func demo3() {
	var slice []string = make([]string, 3, 10)
	slice[0] = "秦始皇"
	slice[1] = "牛顿"
	slice[2] = "阿基米德"
	fmt.Printf("slice 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice, len(slice), cap(slice))
}

// 定义方式三: 使用数组方式定义切片
func demo4() {
	var slice []string = []string{"秦始皇", "牛顿", "三体"}
	fmt.Printf("slice 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice, len(slice), cap(slice))
}

// 切片的遍历
func demo5() {
	var slice []string = []string{"秦始皇", "牛顿", "三体"}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("遍历 slice 元素 = %v \n", slice[i])
	}
	for index, value := range slice {
		fmt.Printf("遍历 slice index = %v, value = %v \n", index, value)
	}
}

// 内置函数 append
func demo6() {
	var slice1 []string = []string{"秦始皇", "牛顿", "面壁者"}
	fmt.Printf("slice1 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice1, len(slice1), cap(slice1))
	// append 本质是底层数组扩容, 如果数组扩容不够(超过了切片的 cap), 则需要重新开辟一块内存存放扩容数据
	slice2 := append(slice1, "三体小说")
	slice2[0] = "始皇大帝"
	// 这里的 slice1、slice2 对应的数组不是同一个(因为 append 超过了 cap)
	fmt.Printf("slice1 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice2, len(slice2), cap(slice2))
	// 如果数组扩容(没超过切片的 cap)够了, 则就在原来的数组上修改, 这里会导致两边的值一起修改
	slice3 := append(slice2, "刘慈欣")
	slice3[0] = "秦始皇大帝"
	// 这里的 slice2、slice3 对应的是同一个数组(因为 append 没有超过 cap)
	fmt.Printf("slice2 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice3, len(slice3), cap(slice3))
}

// 内置函数 copy
func demo7() {
	var slice1 []int32 = []int32{1, 2, 3, 4, 5}
	var slice2 []int32 = make([]int32, 10, 10)
	// 将 slice1 值 copy 到 slice2 中
	copy(slice2, slice1)
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("slice2 = %v\n", slice2)
}

// string 和 slice
func demo8() {
	str := "masteryourself"
	slice := str[10:]
	fmt.Printf("提取字符串 %v \n", slice)
}
