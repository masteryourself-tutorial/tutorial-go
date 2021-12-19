package main

import "fmt"

// map 使用 for-range 遍历
func main() {
	var heroes = map[string]string{
		"ironMan":   "蜘蛛侠",
		"spiderMan": "钢铁侠",
	}
	for k, v := range heroes {
		fmt.Printf("k = %v, v = %v \n", k, v)
	}
}
