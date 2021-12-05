package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Printf("\n==================== 分隔符: demo1 ====================\n")
	demo1()
	fmt.Printf("\n==================== 分隔符: demo2 ====================\n")
	demo2()
	fmt.Printf("\n==================== 分隔符: demo3 ====================\n")
	demo3()
	fmt.Printf("\n==================== 分隔符: demo4 ====================\n")
	demo4()
	fmt.Printf("\n==================== 分隔符: demo5 ====================\n")
	demo5()
	fmt.Printf("\n==================== 分隔符: demo6 ====================\n")
	demo6()
	fmt.Printf("\n==================== 分隔符: demo7 ====================\n")
	demo7()
}

// 创建一个结构体
type Student struct {
	Id   int32
	Name string
	Age  int8
}

// 创建方式一, 在内存中是连续分布的
func demo1() {
	var student Student
	student.Id = 1
	student.Name = "张三"
	student.Age = 18
	fmt.Printf("student = %v \n", student)
	// 这里是连续分布的
	fmt.Printf("内存分布: Id = %p, Name = %p, Age = %p \n", &student.Id, &student.Name, &student.Age)
}

// struc 是值类型
func demo2() {
	var student1 Student
	student1.Id = 1
	student1.Name = "张三"
	student1.Age = 18
	var student2 = student1
	student2.Name = "李四"
	fmt.Printf("student1 = %v \n", student1)
	fmt.Printf("student2 = %v \n", student2)
}

// 创建方式二
func demo3() {
	//如果不赋初始值, 也可以这样定义 var student Student = Student{}
	var student Student = Student{1, "翠芬", 18}
	fmt.Printf("student = %v \n", student)
}

// 创建方式三: 结构体指针 new
func demo4() {
	var student *Student = new(Student)
	// 正常来说, 指针应该这样赋值, 但在 go 中, 也可以这样写
	// student.Id = 1, 底层会自动做转化
	(*student).Id = 1
	student.Name = "威震天"
	student.Age = 99
	fmt.Printf("student = %v \n", *student)
}

// 创建方式四: 结构体指针 &
func demo5() {
	// 也可以这样定义: var student *Student = &Student{1, "擎天柱", 88}
	var student *Student = &Student{}
	student.Name = "擎天柱~"
	fmt.Printf("student = %v \n", *student)
}

// 使用指针修改
func demo6() {
	var student1 Student
	student1.Name = "阿大"
	// 通过指针修改会影响 student1
	var student2 *Student = &student1
	student2.Name = "阿二"
	fmt.Printf("student1 = %v, 地址是 %p \n", student1, &student1)
	fmt.Printf("student2 = %v, 地址是 %p, 存储的值是 %p \n", *student2, &student2, student2)
}

// struct tag 功能, 可以用来做反射处理
type Teacher struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

// json.Marshal() 方法用到了反射
func demo7() {
	teacher := Teacher{18, "仓老师"}
	res, err := json.Marshal(teacher)
	if err != nil {
		fmt.Println("json 格式转换错误", err)
	}
	fmt.Printf("json = %v \n", string(res))
}
