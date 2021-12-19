package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func main() {
	stu := Stu{18, "tom"}
	res, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json 格式转换错误", err)
	}
	fmt.Printf("json = %v \n", string(res))
}
