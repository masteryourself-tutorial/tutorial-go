package main

import "fmt"

func main() {
	mapDefine1()
	mapDefine2()
	mapDefine3()
}

// 定义方式 1: 先定义, 再 make()
func mapDefine1() {
	// map 在使用前一定要先定义
	var heroes map[string]string
	heroes = make(map[string]string, 8)
	heroes["ironMan"] = "蜘蛛侠"
	heroes["spiderMan"] = "钢铁侠"
	fmt.Printf("heroes = %v \n", heroes)
}

// 定义方式 2: 定义并 make()
func mapDefine2() {
	var heroes = make(map[string]string, 8)
	heroes["ironMan"] = "蜘蛛侠"
	heroes["spiderMan"] = "钢铁侠"
	fmt.Printf("heroes = %v \n", heroes)
}

// 定义方式 3: 定义并直接赋值
func mapDefine3() {
	var heroes = map[string]string{
		"ironMan":   "蜘蛛侠",
		"spiderMan": "钢铁侠",
	}
	fmt.Printf("heroes = %v \n", heroes)
}
