package main

import (
	"fmt"
	"go-core/07_struct/06_factory/model"
)

func main() {
	stu := model.NewStu(10, "张三")
	fmt.Println("stu = ", stu)
}
