package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonMarshal()
	jsonUnmarshal()
}

// json 序列化
func jsonMarshal() {
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
	// hero 序列化 {"id":18,"name":"迪迦","address":["地球","太阳系"],"boy":true}
	fmt.Printf("hero 序列化 %v \n", string(heroJsonStr))
}

// json 反序列化
func jsonUnmarshal() {
	jsonStr := "{\"id\":18,\"name\":\"迪迦\",\"address\":[\"地球\",\"太阳系\"],\"boy\":true}"
	var heroJson Hero
	err := json.Unmarshal([]byte(jsonStr), &heroJson)
	if err != nil {
		fmt.Println("json 反序列化失败 \n", err)
	}
	// {18 迪迦 [地球 太阳系] true}
	fmt.Println(heroJson)
}

type Hero struct {
	Id      int32    `json:"id"`
	Name    string   `json:"name"`
	Address []string `json:"address"`
	Boy     bool     `json:"boy"`
}
