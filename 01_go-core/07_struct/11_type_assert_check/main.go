package main

import "fmt"

func main() {
	var x interface{}
	stu := Stu{10}
	// 把 stu 赋值给空接口 x
	x = stu
	// 类型断言 + 检测
	if z, ok := x.(Stu); ok {
		fmt.Printf("z 类型是 %T \n", z)
		fmt.Println("stu id = ", z.GetId())
	}
}

type Stu struct {
	id int32
}

func (s Stu) GetId() int32 {
	return s.id
}
