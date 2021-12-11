package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var hero = Hero{
		Id:      18,
		Name:    "迪迦",
		Address: []string{"地球", "太阳系"},
		Boy:     true,
	}
	heroJsonStr, err := json.Marshal(hero)
	if err != nil {
		fmt.Println("json 序列化失败", err)
	}
	fmt.Printf("hero 序列化 %v \n", string(heroJsonStr))
	var heroJson Hero
	err = json.Unmarshal(heroJsonStr, &heroJson)
	if err != nil {
		fmt.Println("json 反序列化失败 \n", err)
	}
	fmt.Printf("hero 反序列化 %v \n", heroJson)
}

type Hero struct {
	Id      int32    `json:"id"`
	Name    string   `json:"name"`
	Address []string `json:"address"`
	Boy     bool     `json:"boy"`
}
