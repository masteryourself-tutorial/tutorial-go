package model

type student struct {
	Id   int
	name string
}

func NewStudent() student {
	return student{}
}

func (student *student) GetName() string {
	return student.name
}

func (student *student) SetName(name string) {
	student.name = name
}
