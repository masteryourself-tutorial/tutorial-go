package main

import "fmt"

type Stu struct {
	Id   int
	Age  int16
	Name string
}

func main() {
	stu := Stu{}
	// stu 的地址是 0xc000052020，id = 0xc000052020, age = 0xc000052028, name = 0xc000052030
	fmt.Printf("stu 的地址是 %p，id = %p, age = %p, name = %p \n",
		&stu, &stu.Id, &stu.Age, &stu.Name)
}
