package model

import "fmt"

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机结束工作")
}
