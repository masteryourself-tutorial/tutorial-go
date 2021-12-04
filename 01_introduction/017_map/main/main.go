package main

import (
	"fmt"
)

func main()  {
	fmt.Printf("\n==================== 分隔符: demo1 ====================\n")
	demo1()
	fmt.Printf("\n==================== 分隔符: demo2 ====================\n")
	demo2()
	fmt.Printf("\n==================== 分隔符: demo3 ====================\n")
	demo3()
	fmt.Printf("\n==================== 分隔符: demo4 ====================\n")
	demo4()
	fmt.Printf("\n==================== 分隔符: demo5 ====================\n")
	demo5()
}

// 定义方式 1: 先定义, 再 make()
func demo1()  {
	// map 在使用前一定要先定义
	// map 的 key 是不可重复的, 无序的
	var heroes map[string]string 
	heroes = make(map[string]string, 8)
	heroes["ironMan"] = "蜘蛛侠"
	heroes["spiderMan"] = "钢铁侠"
	fmt.Printf("heroes = %v \n", heroes)
}

// 定义方式 2: 定义并 make()
func demo2()  {
	var heroes map[string]string  = make(map[string]string, 8)
	heroes["ironMan"] = "蜘蛛侠"
	heroes["spiderMan"] = "钢铁侠"
	fmt.Printf("heroes = %v \n", heroes)
}

// 定义方式 3: 定义并直接赋值
func demo3()  {
	var heroes map[string]string  = map[string]string{
		"ironMan" : "蜘蛛侠",
		"spiderMan" : "钢铁侠",
	}
	fmt.Printf("heroes = %v \n", heroes)
}

// 增删改查
func demo4()  {
	var heroes map[string]string  = make(map[string]string, 8)
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
	var name string = "monkey"
	hero, exist := heroes[name]
	if exist {
		fmt.Printf("查找(找到): name = %v, hero = %v \n", name, hero)
	} else {
		fmt.Printf("查找(未找到): name = %v, \n", name)
	}
}

// map 的遍历
func demo5()  {
	var heroes map[string]string  = make(map[string]string, 8)
	heroes["ironMan"] = "蜘蛛侠"
	heroes["spiderMan"] = "钢铁侠"
	for name, hero := range heroes {
		fmt.Printf("name = %v, hero = %v \n", name, hero)
	}
}