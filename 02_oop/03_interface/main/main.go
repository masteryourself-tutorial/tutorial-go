package main

import (
	"fmt"
	"sort"
	"tutorial.masteryourself.pers/tutorial_go/02_oop/03_interface/hero"
	"tutorial.masteryourself.pers/tutorial_go/02_oop/03_interface/model"
)

func main() {
	fmt.Printf("\n==================== 分隔符: demo1 ====================\n")
	demo1()
	fmt.Printf("\n==================== 分隔符: demo2 ====================\n")
	demo2()
}

func demo1() {
	camera := model.Camera{}
	phone := model.Phone{}
	computer := model.Computer{}
	computer.Work(camera)
	computer.Work(phone)
}

func demo2() {
	hero1 := hero.Hero{Id: 1, Name: "孙悟空"}
	hero2 := hero.Hero{Id: 2, Name: "猪悟能"}
	hero3 := hero.Hero{Id: 3, Name: "沙悟僧"}
	var hs hero.HeroSlice
	hs = append(hs, hero2, hero3, hero1)
	fmt.Println("排序前", hs)
	sort.Sort(hs)
	fmt.Println("排序后", hs)
}
