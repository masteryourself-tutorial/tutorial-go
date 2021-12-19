package main

import "fmt"

func main() {
	switchDemo1()
	switchDemo2()
	switchDemo3()
}

// switch/case 后面跟的是表达式
func switchDemo1() {
	var key = 6
	switch key + 1 {
	case 6:
		fmt.Println("a")
	case 2 + 5:
		fmt.Println("7")
	}
}

// switch 后也可以不跟表达式，类似 if-else 分支
func switchDemo2() {
	var score = 90
	switch {
	case score < 60:
		fmt.Println("不及格")
	case score < 90:
		fmt.Println("良好")
	case score < 100:
		fmt.Println("优秀")
	default:
		fmt.Println("你考了满分了")
	}
}

// fallthrough 穿透
func switchDemo3() {
	var score = 95
	switch {
	case score > 60:
		fmt.Println("及格")
		fallthrough
	case score > 90:
		fmt.Println("良好")
	default:
		fmt.Println("未及格")
	}
}
