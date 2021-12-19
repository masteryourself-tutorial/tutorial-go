package main

import "fmt"

func main() {
	strB := StructB{}
	// 必须要指定结构体名称 strB.Id 会报错
	fmt.Println(strB.StrA.Id)
	fmt.Println(strB.Name)
}

type StructA struct {
	Id int32
}

type StructB struct {
	// 嵌套了有名结构体，变成了组合关系
	StrA StructA
	Name string
}
