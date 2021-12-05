package main

import (
	"fmt"
	"tutorial.masteryourself.pers/tutorial_go/02_oop/01_encapsulation/model"
)

func main() {
	student := model.NewStudent()
	student.Id = 18
	student.SetName("擎天柱")
	fmt.Printf("学生信息是: id = %v, name = %v \n", student.Id, student.GetName())
}
