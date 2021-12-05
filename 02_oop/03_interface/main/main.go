package main

import "tutorial.masteryourself.pers/tutorial_go/02_oop/03_interface/model"

func main() {
	camera := model.Camera{}
	phone := model.Phone{}
	computer := model.Computer{}
	computer.Work(camera)
	computer.Work(phone)
}
