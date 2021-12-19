package main

import "fmt"

func main() {
	r1, r2 := replaceValue(10, 20)
	fmt.Printf("res1 = %v, res2 = %v \n", r1, r2)
}

// 定义时对返回值命名
func replaceValue(x1 int, x2 int) (res1 int, res2 int) {
	res1, res2 = x2, x1
	return res1, res2
}
