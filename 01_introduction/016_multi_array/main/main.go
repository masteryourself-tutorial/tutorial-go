package main

import (
	"fmt"
)

func main()  {
	fmt.Printf("\n==================== 分隔符: demo1 ====================\n")
	demo1()
	fmt.Printf("\n==================== 分隔符: demo2 ====================\n")
	demo2()
	fmt.Printf("\n==================== 分隔符: demo3 ====================\n")
	demo3()
}

// 二维数组定义方式 1
func demo1()  {
	// 三行四列
	var multiArray [3][4]int
	multiArray[1][1] = 10
	for i := 0; i < len(multiArray); i++ {
		for j := 0; j < len(multiArray[i]); j++ {
			fmt.Printf("%v \t", multiArray[i][j])
		}
		fmt.Println()
	}
}

// 二维数组地址值
func demo2()  {
	var multiArray [3][4]int
	fmt.Printf("二位数组地址是 %p \n", &multiArray)
	// 它们之间相差 32 个字节, 这是因为每行有四个元素, 一共占用 4 * 8 = 32 字节
	fmt.Printf("二位数组第一排元素首地址是 %p \n", &multiArray[0])
	fmt.Printf("二位数组第二排元素首地址是 %p \n", &multiArray[1])
	// 它们之间相差 8 个字节, 这是因为每个元素是 8 个字节
	fmt.Printf("二位数组第一排第一个元素首地址是 %p \n", &multiArray[0][0])
	fmt.Printf("二位数组第一排第二个元素首地址是 %p \n", &multiArray[0][1])
}

// 二维数组定义方式 2
func demo3()  {
	var multiArray [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("multiArray = %v \n", multiArray)
}