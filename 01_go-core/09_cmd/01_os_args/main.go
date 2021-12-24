package main

import (
	"fmt"
	"os"
)

// 以空格分隔
func main() {
	args := os.Args
	fmt.Println("args = ", args)
	for i, arg := range args {
		fmt.Printf("index = %v, arg = %v \n", i, arg)
	}
}
