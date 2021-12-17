package main

import (
	"fmt"
	"sync"
	"time"
)

var factMap = make(map[int]int)
var lock = sync.Mutex{}

func main() {
	for i := 1; i <= 20; i++ {
		go fact(i)
	}
	time.Sleep(time.Second * 2)
	for index, value := range factMap {
		fmt.Printf("factMap[%v] = %v \n", index, value)
	}
}

func fact(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	// 这里会产生并发问题, 使用 Mutex 上锁
	lock.Lock()
	factMap[n] = res
	lock.Unlock()
}
