package main

import "fmt"

// 一次性申明多个全局变量
var (
	X1 = 1000
	X2 = 2000
	X3 = 3000
)

func main() {
	// 声明多个变量方式 1
	var x1, x2, x3 int
	fmt.Printf("x1 = %v, x2 = %v, x3 = %v \n", x1, x2, x3)

	// 声明多个变量方式 2
	var y1, y2, y3 = 22, "tom", 33
	fmt.Printf("y1 = %v, y2 = %v, y3 = %v \n", y1, y2, y3)

	// 声明多个变量方式 3
	z1, z2, z3 := 222, "jack", 333
	fmt.Printf("z1 = %v, z2 = %v, z3 = %v \n", z1, z2, z3)

	// 打印全局变量
	fmt.Printf("X1 = %v, X2 = %v, X3 = %v \n", X1, X2, X3)
}
