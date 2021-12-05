package model

// 这里定义的是小写, 所以其他包不可访问
type student struct {
	Id   int
	Name string
}

// 通过工厂模式对外提供一个公共方法
func NewStudent(id int, name string) student {
	return student{
		Id:   id,
		Name: name,
	}
}
