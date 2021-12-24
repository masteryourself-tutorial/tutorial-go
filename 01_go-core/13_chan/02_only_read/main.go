package main

import "fmt"

// 管道可以看做成队列, FIFO
func main() {
	onlyReadChan()
}

func onlyReadChan() {
	var ch = make(chan int, 3)
	ch <- 2
	ch <- 5
	ch <- 8
	close(ch)
	receive(ch)
}

// 接收一个只读 chan, 这样就可以防止误操作了
func receive(ch <-chan int) {
	for v := range ch {
		fmt.Printf("从 chan 中读取数据 %v \n", v)
	}
}
