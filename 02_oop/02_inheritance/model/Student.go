package model

import "fmt"

type Student struct {
	Age   int16
	Name  string
	score int16
}

func (student *Student) SetScore(score int16) {
	student.score = score
}

func (student Student) Show() {
	fmt.Printf("学生信息是 %v \n", student)
}
