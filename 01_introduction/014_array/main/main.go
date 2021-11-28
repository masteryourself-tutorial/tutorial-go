package main

import (
	"fmt"
)

func main()  {
	fmt.Println("==================== 分隔符: demo1 ====================")
	demo1()
	fmt.Println("==================== 分隔符: demo2 ====================")
	demo2()
	fmt.Println("==================== 分隔符: demo3 ====================")
	demo3()
	fmt.Println("==================== 分隔符: demo4 ====================")
	demo4()
	fmt.Println("==================== 分隔符: demo5 ====================")
	var arr [3]int = [3]int {1, 2, 3}
	demo5(arr)
	fmt.Printf("main 中的 array 是 %v \n", arr)
	fmt.Println("==================== 分隔符: demo6 ====================")
	demo6(&arr)
	fmt.Printf("main 中的 array 是 %v \n", arr)
}

func demo1()  {
	var hens[3] float64
	hens[0] = 2.6
	hens[1] = 3.8
	hens[2] = 5.3
	totalWeigt := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeigt += hens[i]
	}
	avgWeight := totalWeigt / float64(len(hens))
	fmt.Printf("鸡总体重 %v, 平均体重 %.2f \n", totalWeigt, avgWeight)
}

func demo2()  {
	// 64 位下, 每个元素占用 8 个字节
	var intArr [3]int
	fmt.Printf("数组地址是 %p \n", &intArr)
	// 数组地址和第 1 个元素地址值是相等的
	fmt.Printf("数组中第 1 个元素的地址是 %p \n", &intArr[0])
	// 因为数组是连续的, 所以后面元素的地址就是 [第 1 个元素的地址值 + 8]
	fmt.Printf("数组中第 2 个元素的地址是 %p \n", &intArr[1])
	fmt.Printf("数组中第 3 个元素的地址是 %p \n", &intArr[2])
}

func demo3()  {
	var arr1 [3]int= [3]int {1, 2, 3}
	fmt.Printf("arr1 = %v \n", arr1)

	var arr2 = [3]int {4, 5, 6}
	fmt.Printf("arr2 = %v \n", arr2)
	
	var arr3 = [...]int {7, 8, 9}
	fmt.Printf("arr3 = %v \n", arr3)
	
	var arr4 = [...]int {0:10, 2:11}
	fmt.Printf("arr4 = %v \n", arr4)
}

func demo4()  {
	var heros [3]string = [3]string {"蜘蛛侠", "钢铁侠", "绿箭侠"}
	for index, value := range heros{
		fmt.Printf("index = %v, value = %v \n", index, value)
	}
}

func demo5(arr [3]int)  {
	arr[0] = 998
	fmt.Printf("demo5 中的 array 是 %v , 说明数组是值传递, 属于值类型\n", arr)
}

func demo6(arr *[3]int)  {
	(*arr)[0] = 998
	fmt.Printf("demo5 中的 array 是 %v , 数组可以通过指针修改值\n", *arr)
}