package main

import "fmt"

// 增删改查
func main() {
	var heroes = make(map[string]string, 8)
	// 增加
	heroes["ironMan"] = "蜘蛛侠"
	heroes["spiderMan"] = "钢铁侠"
	heroes["monkey"] = "孙悟空"
	fmt.Printf("增加: heroes = %v \n", heroes)
	// 删除
	delete(heroes, "ironMan")
	fmt.Printf("删除: heroes = %v \n", heroes)
	// 修改
	heroes["monkey"] = "齐天大圣"
	fmt.Printf("修改: heroes = %v \n", heroes)
	// 查询
	var name = "monkey"
	hero, exist := heroes[name]
	if exist {
		fmt.Printf("查找(找到): name = %v, hero = %v \n", name, hero)
	} else {
		fmt.Printf("查找(未找到): name = %v, \n", name)
	}
}
