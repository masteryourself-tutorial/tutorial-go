package dao

import (
	"chat/model"
	"chat/utils"
	"fmt"
)

func GetUserList() []*model.User {
	db := utils.GetDb()
	var data = make([]*model.User, 10)
	db.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
