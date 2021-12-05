package main

import "tutorial.masteryourself.pers/tutorial_go/02_oop/02_inheritance/model"

func main() {
	// 继承写法
	pupil := model.Pupil{}
	pupil.Student.Name = "翠芬"
	pupil.Student.Age = 8
	pupil.Student.SetScore(95)
	pupil.Testing()
	pupil.Student.Show()
	// go 语言可以简写成这样, 在原方法中找不到这个对象, 就去匿名结构体中查找, 遵循就近访问原则
	graduate := model.Graduate{}
	graduate.Name = "翠花"
	graduate.Age = 18
	graduate.SetScore(65)
	graduate.Testing()
	graduate.Show()
}
