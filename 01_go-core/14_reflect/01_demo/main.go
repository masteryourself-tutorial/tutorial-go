package main

import (
	"fmt"
	"reflect"
)

func main() {
	stu := Stu{1, "张三"}
	reflectTest(stu)
}

func reflectTest(i interface{}) {
	// 1. 通过反射获取 type
	rType := reflect.TypeOf(i)
	fmt.Printf("rType = %T \n", rType)
	// 2. 通过反射获取 value
	rValue := reflect.ValueOf(i)
	fmt.Printf("rValue = %v \n", rValue)
	// 3. 将 rValue 转成 interface{}
	iInterface := rValue.Interface()
	fmt.Printf("iInterface = %T \n", iInterface)
	// 4. 使用类型断言转成需要的类型
	iValue := iInterface.(Stu)
	fmt.Printf("iValue = %v \n", iValue)
}

type Stu struct {
	Id   int
	Name string
}
