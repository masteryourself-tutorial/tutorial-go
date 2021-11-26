package main

import (
	"fmt"
)

func main()  {

	var name string
	var age int32
	var salary float32
	var isPass bool

	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	fmt.Println("请依次输入年龄, 工资, 是否通过考试, 并以空格分隔")
	fmt.Scanf("%d %f %t", &age, &salary, &isPass)

	fmt.Printf("员工信息: 姓名 %v, 年龄 %v, 工资 %v, 是否通过考试 %v", name, age, salary, isPass)
	
}