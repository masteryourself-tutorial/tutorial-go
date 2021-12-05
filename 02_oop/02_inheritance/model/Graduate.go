package model

import "fmt"

type Graduate struct {
	Student
}

func (g Graduate) Testing() {
	fmt.Println("大学生考试中")
}
