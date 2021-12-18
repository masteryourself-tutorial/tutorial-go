package main

import "fmt"

func main() {
	variableDefine1()
	variableDefine2()
	variableDefine3()
}

// 变量声明方式 1：指定变量类型，声明后若不赋值，使用默认值
func variableDefine1() {
	var i int
	fmt.Println("i = ", i)
}

// 变量声明方式 2：根据值类型自动推导变量类型
func variableDefine2() {
	var i = 22
	fmt.Println("i = ", i)
}

// 变量声明方式 3. 省略 var，使用 :=
func variableDefine3() {
	i := 33
	fmt.Println("i = ", i)
}
