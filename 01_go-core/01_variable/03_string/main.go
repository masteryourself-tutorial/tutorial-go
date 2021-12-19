package main

import "fmt"

func main() {
	// \t 会被转义
	str1 := "123\t456"
	// \t 会原样输出
	str2 := `123\t456`
	fmt.Printf("str1 = %v \n", str1)
	fmt.Printf("str2 = %v \n", str2)
}
