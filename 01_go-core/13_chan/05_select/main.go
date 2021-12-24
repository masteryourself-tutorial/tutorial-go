package main

import "fmt"

func main() {
	var ch = make(chan int, 3)
	ch <- 2
	ch <- 5
	ch <- 8
	selectChan(ch)
}

func selectChan(ch chan int) {
exitFlag:
	for {
		select {
		case v := <-ch:
			fmt.Println("从 chan 中获取数据", v)
		default:
			fmt.Println("获取不到数据, 终止循环操作")
			break exitFlag
		}
	}
}
