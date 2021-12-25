package main

import (
	"fmt"
	"reflect"
)

func main() {
	stu := Stu{1, "张三"}
	reflectTest(stu)
	fmt.Println("stu = ", stu)
}

func reflectTest(i interface{}) {
	// 1. 通过反射获取 value
	rValue := reflect.ValueOf(i)
	fmt.Printf("rValue = %T \n", rValue)
	fieldNum := rValue.NumField()
	fmt.Printf("传入类型有 %d 个字段 \n", fieldNum)
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("第 %d 个方法名是 %v \n", i, rValue.Field(i).Type())
	}
	rValue.Elem().MethodByName("Print").Call(nil)
}

type Stu struct {
	Id   int
	Name string
}

func (s Stu) Print() {
	fmt.Println("print go", s)
}
