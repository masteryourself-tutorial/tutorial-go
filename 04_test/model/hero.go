package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Hero struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Skill []string `json:"skill"`
}

var fileName = "/Users/ruanrenzhao/go_project/tutorial-go/04_test/doc/json.txt"

func (h *Hero) Store() {
	// 序列化成 heroJson 字符串
	data, err := json.Marshal(h)
	if err != nil {
		fmt.Println("json 序列化失败", err)
		return
	}
	// 写入文件
	err = ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		fmt.Println("文件写入失败", err)
	}
}

func (h *Hero) ReStore() {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	err = json.Unmarshal(data, h)
	if err != nil {
		fmt.Println("json 反序列化", err)
	}
}
