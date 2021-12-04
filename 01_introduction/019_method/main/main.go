package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n==================== 分隔符: demo1 ====================\n")
	demo1()
	fmt.Printf("\n==================== 分隔符: demo2 ====================\n")
	demo2()
	// fmt.Printf("\n==================== 分隔符: demo3 ====================\n")
	// demo3()
	// fmt.Printf("\n==================== 分隔符: demo4 ====================\n")
	// demo4()
	// fmt.Printf("\n==================== 分隔符: demo5 ====================\n")
	// demo5()
	// fmt.Printf("\n==================== 分隔符: demo6 ====================\n")
	// demo6()
	// fmt.Printf("\n==================== 分隔符: demo7 ====================\n")
	// demo7()
}

// 定义一个结构体 Person
type Person struct {
	Id int
	Name string
}

// 给 person 绑定一个方法
func (person Person) method1() {
	fmt.Printf("执行 Person method1() 方法: id = %v, name = %v \n", person.Id, person.Name)
}

// 为了提高效率, 通常使用指针绑定
func (person *Person) method2() {
	// 理论上应该这样写才正确, 但 go 语言为了方便, 也可以直接简写
	fmt.Printf("执行 Person method2() 方法: id = %v, name = %v \n", (*person).Id, (*person).Name)
	fmt.Printf("执行 Person method2() 方法: id = %v, name = %v \n", person.Id, person.Name)
}

// 调用方法和调用函数机制非常类似, 只不过调用方法会把当前变量本身也传递到方法中去, 类比 Java 中的 this
// 如果变量是值类型, 则进行值拷贝, 如果是引用类型, 则进行地址拷贝
func demo1()  {
	person := Person{1, "亚古兽"}
	person.method1()
}

// 调用方法和调用函数机制非常类似, 只不过调用方法会把当前变量本身也传递到方法中去, 类比 Java 中的 this
// 如果变量是值类型, 则进行值拷贝, 如果是引用类型, 则进行地址拷贝
func demo2()  {
	person := Person{2, "加鲁鲁兽"}
	// 理论上应该这样调用才正确, go 语言中为了方便使用, 可以直接这样写
	// (&person).method2()
	person.method2()
}