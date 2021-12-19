package main

import "fmt"

func main() {
	var stu = Stu{1, "tom"}
	fmt.Printf("使用值类型调用 %v \n", stu.getId())
	fmt.Printf("使用值类型调用 %v \n", (&stu).getName())
	fmt.Printf("使用引用类型调用 %v \n", stu.getId())
	fmt.Printf("使用引用类型调用 %v \n", (&stu).getName())
}

type Stu struct {
	id   int32
	name string
}

// 接收参数为值类型
func (s Stu) getName() string {
	return s.name
}

// 接收参数为引用类型
func (s *Stu) getId() int32 {
	return (*s).id
}
