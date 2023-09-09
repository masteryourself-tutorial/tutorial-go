package service

import (
	"chat/dao"
	"chat/model"
)

func GetUserList() []*model.User {
	return dao.GetUserList()
}
