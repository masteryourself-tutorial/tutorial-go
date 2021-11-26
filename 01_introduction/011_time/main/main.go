package main

import (
	"fmt"
	"time"
)

func main()  {
	// 获取当前时间
	now := time.Now()
	fmt.Println("当前时间 \n", now)

	// 格式化日期, 2006-01-02 15:04:05 是固定的
	fmt.Printf("日期格式化方式1 %d-%d-%d %d:%d:%d \n", 
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("日期格式化方式2 %v \n", now.Format("2006-01-02 15:04:05"))

	fmt.Printf("当前年月日 %v \n", now.Format("2006/01/02"))

}