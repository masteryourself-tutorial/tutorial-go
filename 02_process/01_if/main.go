package main

import "fmt"

func main() {
	ifDemo1()
	ifDemo2()
}

// if 后面跟条件表达式, 大括号不能省略, else 也不能换行
func ifDemo1() {
	var age = 19
	if age > 18 {
		fmt.Println("你成年了，可以去网吧")
	} else {
		fmt.Println("你未成年，不能去网吧")
	}
}

// 支持在 if 中直接定义一个变量
func ifDemo2() {
	if age := 19; age > 18 {
		fmt.Println("你成年了，可以去网吧")
	} else {
		fmt.Println("你未成年，不能去网吧")
	}
}
