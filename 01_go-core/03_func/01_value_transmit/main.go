package main

import "fmt"

func main() {
	var i = 10
	modify(i)
	// main() 中的 i =  10
	fmt.Println("main() 中的 i = ", i)
}

// 值类型不会影响原来的值
func modify(i int) {
	i = i + 10
	// main() 中的 i =  10
	fmt.Println("modify() 中的 i = ", i)
}
