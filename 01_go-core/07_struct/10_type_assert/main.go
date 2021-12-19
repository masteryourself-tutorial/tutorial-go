package main

import "fmt"

func main() {
	var x interface{}
	stu := Stu{10}
	// 把 stu 赋值给空接口 x
	x = stu
	// 如果想通过 x 访问 GetId 方法, 就需要使用类型断言
	fmt.Println("stu id = ", x.(Stu).GetId())
}

type Stu struct {
	id int32
}

func (s Stu) GetId() int32 {
	return s.id
}
