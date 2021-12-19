package main

import "fmt"

type Stu struct {
	id int
}

// 结构体是值类型, 不影响之前的值
func main() {
	var stu1 = Stu{1}
	var stu2 = stu1
	stu2.id = 10
	// stu1 =  {1}
	fmt.Println("stu1 = ", stu1)
	// stu2 =  {10}
	fmt.Println("stu2 = ", stu2)
}
