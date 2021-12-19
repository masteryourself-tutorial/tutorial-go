package main

import "fmt"

func main() {
	multiArrayDefine1()
	multiArrayDefine2()
}

// 定义方式1：先声明再赋值
func multiArrayDefine1() {
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

// 定义方式2：直接初始化
func multiArrayDefine2() {
	var multiArray = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("multiArray = %v \n", multiArray)
}
