package model

// 定义一个结构体, 类名小写
type stu struct {
	Id   int32
	Name string
}

// 因为 stu 是小写的只能在 model 包中访问
// 所以提供一个工厂方法来获取
func NewStu(id int32, name string) *stu {
	return &stu{
		Id:   id,
		Name: name,
	}
}
