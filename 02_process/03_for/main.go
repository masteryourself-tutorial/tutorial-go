package main

import "fmt"

func main() {
	forDemo1()
	forDemo2()
	forDemo3()
	forDemo4()
}

func forDemo1() {
	for i := 0; i < 3; i++ {
		fmt.Println("蜘蛛侠")
	}
}

func forDemo2() {
	i := 0
	for i < 3 {
		fmt.Println("蝙蝠侠")
		i++
	}
}

func forDemo3() {
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("钢铁侠")
		i++
	}
}

func forDemo4() {
	str := "hello world"
	for i, v := range str {
		fmt.Printf("index = %v, value = %c \n", i, v)
	}
}
