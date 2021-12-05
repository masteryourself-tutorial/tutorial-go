package model

import "fmt"

type Pupil struct {
	Student
}

func (p Pupil) Testing() {
	fmt.Println("小学生考试中")
}
