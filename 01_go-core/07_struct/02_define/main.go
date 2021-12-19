package main

import "fmt"

func main() {
	// 定义方式 1
	var stu1 Stu
	// stu1 = {0 }
	fmt.Printf("stu1 = %v \n", stu1)

	// 定义方式 2
	var stu2 Stu = Stu{10, "tom"}
	// stu2 = {10 tom}
	fmt.Printf("stu2 = %v \n", stu2)

	// 定义方式 3：这种方式返回的是结构体指针
	var stu3 *Stu = new(Stu)
	// 因为 stu3 是一个指针，所以应该这样赋值
	// go 为了简化开发也可以这样使用 stu3.id = 20
	(*stu3).id = 20
	stu3.Name = "lucy"
	// stu3 = &{20 lucy}
	fmt.Printf("stu3 = %v \n", stu3)

	// 定义方式 4：这种方式返回的是结构体指针
	var stu4 *Stu = &Stu{}
	stu4.id = 30
	(*stu4).Name = "Jack"
	// stu4 = &{30 Jack}
	fmt.Printf("stu4 = %v \n", stu4)
}

type Stu struct {
	id   int
	Name string
}
