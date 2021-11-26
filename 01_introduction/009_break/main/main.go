package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100)
		if num == 99 {
			break
		}
		count++
	}
	fmt.Printf("随机生成 99 耗费了 %v 次", count)
}