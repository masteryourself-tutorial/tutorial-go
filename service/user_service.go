package service

import (
	"chat/dao"
	"chat/model"
)

func SelectUser(id uint) model.User {
	return dao.SelectUser(id)
}

func CreateUser(user model.User) {
	dao.CreateUser(user)
}

func UpdateUser(user model.User) {
	dao.UpdateUser(user)
}

func DeleteUser(user model.User) {
	dao.DeleteUser(user)
}

func GetUserList() []model.User {
	return dao.GetUserList()
}
