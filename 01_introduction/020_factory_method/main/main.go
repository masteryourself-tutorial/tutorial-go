package main

import (
	"fmt"
	"tutorial.masteryourself.pers/tutorial_go/01_introduction/020_factory_method/model"
)

func main() {
	var student = model.NewStudent(1, "钢铁侠")
	fmt.Printf("student = %v \n", student)
}
