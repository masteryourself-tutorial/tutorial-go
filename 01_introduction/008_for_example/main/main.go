package main

import (
	"fmt"
)

func main()  {
	demo1()
}

func demo1()  {
	cycyleCount := 9
	for i := 1; i <= cycyleCount; i++ {
		lineNumber := 2 * i - 1
		// 打印空格
		for j := 0; j < cycyleCount - i; j++ {
			fmt.Printf(" ")
		}
		// 打印星星
		for j := 0; j < lineNumber; j++ {
			// 判断是否打星
			if j == 0 || j == lineNumber - 1 || i == cycyleCount {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		// 打印下一行
		fmt.Println()
	}
}