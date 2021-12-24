package main

import "fmt"

// 管道可以看做成队列, FIFO
func main() {
	channelDemo()
	channelLoop()
}

// chanel demo
func channelDemo() {
	// channel 是引用类型, 必须初始化后才能使用, 定义管道容量后, 不能超出
	var intChan chan int = make(chan int, 3)
	// 由此可见 intChan 存储的也是个地址, 指向的内存中的某个 channel
	fmt.Printf("intChan 的地址是 %v, 存储的值是 %v \n", &intChan, intChan)
	// 向管道中存入数据
	intChan <- 10
	intChan <- 20
	fmt.Printf("intChan 当前存储数据 = %v, 总容量 =%v \n", len(intChan), cap(intChan))
	intChan <- 30
	// 从管道中取出数据
	num1 := <-intChan
	num2 := <-intChan
	fmt.Printf("intChan 当前存储数据 = %v, 总容量 =%v \n", len(intChan), cap(intChan))
	fmt.Printf("num1 = %v, num2 = %v \n", num1, num2)
}

// channel 遍历
func channelLoop() {
	var intChan = make(chan int, 3)
	for i := 0; i < cap(intChan); i++ {
		intChan <- i * 100
	}
	// 要遍历 channel, 必须要调用 close() 方法
	close(intChan)
	for i := range intChan {
		fmt.Printf("从 intChan 中获取数据 %v \n", i)
	}
}
