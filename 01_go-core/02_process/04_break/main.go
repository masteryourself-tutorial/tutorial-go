package main

import "fmt"

// 定义一个结束标签来指明语句块
func main() {
closeFlag:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break closeFlag
			}
			fmt.Println("j = ", j)
		}
	}
}
