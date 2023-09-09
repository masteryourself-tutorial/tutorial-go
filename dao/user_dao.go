package dao

import (
	"chat/model"
	"chat/utils"
)

func SelectUser(id uint) (user model.User) {
	db := utils.GetDb()
	db.First(&user, "id = ?", id)
	return user
}

func CreateUser(user model.User) {
	db := utils.GetDb()
	db.Create(&user)
}

func UpdateUser(user model.User) {
	db := utils.GetDb()
	db.Updates(&user)
}

func DeleteUser(user model.User) {
	db := utils.GetDb()
	db.Delete(&user)
}

func GetUserList() []model.User {
	db := utils.GetDb()
	var data = make([]model.User, 10)
	db.Find(&data)
	return data
}
