package main

import "fmt"

func main() {
	monkey := Monkey{}
	monkey.Eat()
	// 接口不能创建实例, 但可以指向一个实现该接口的自定义类型的变量
	var fly Fly = monkey
	fly.Fly()
}

// 接口飞翔
type Fly interface {
	Fly()
}

// 接口吃饭
type Eat interface {
	Eat()
}

// 孙悟空实现类，实现了多个接口
type Monkey struct {
}

func (m Monkey) Fly() {
	fmt.Println("猴子学会飞翔")
}

func (m Monkey) Eat() {
	fmt.Println("猴子学会吃东西")
}
